package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// ตัวแปร db ประกาศไว้ตรงนี้เพื่อให้ฟังก์ชันอื่นๆ เรียกใช้งานฐานข้อมูลได้
var db *sql.DB

type SecurityLog struct {
	EventType      string `json:"event_type"`
	SourceIP       string `json:"source_ip"`
	TargetEndpoint string `json:"target_endpoint"`
	Severity       string `json:"severity"`
	Description    string `json:"description"`
}

type SecurityLogResponse struct {
	ID             int    `json:"id"`
	EventType      string `json:"event_type"`
	SourceIP       string `json:"source_ip"`
	TargetEndpoint string `json:"target_endpoint"`
	Severity       string `json:"severity"`
	Description    string `json:"description"`
	CreatedAt      string `json:"created_at"`
}

// ฟังก์ชัน API ดึงข้อมูล GET (รวมระบบ API Key แล้ว)
func getLogsHandler(w http.ResponseWriter, r *http.Request) {
	// 1. ดักจับ API Key ก่อนเลย
	apiKey := r.Header.Get("X-API-Key")
	if apiKey != "secret-token-2026" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"error": "Unauthorized - รหัส API Key ไม่ถูกต้องหรือไม่ได้แนบมา"}`)
		return
	}

	// 2. ดักว่าต้องเป็น Method GET เท่านั้น
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 3. สั่ง SQL ดึงข้อมูลทั้งหมด
	rows, err := db.Query("SELECT id, event_type, source_ip, target_endpoint, severity, description, created_at FROM security_logs ORDER BY created_at DESC")
	if err != nil {
		log.Println("ดึงข้อมูลพลาด:", err)
		http.Error(w, "เซิร์ฟเวอร์ขัดข้อง", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// 4. วนลูปจับข้อมูลใส่กล่อง
	var logs []SecurityLogResponse
	for rows.Next() {
		var logItem SecurityLogResponse
		err := rows.Scan(&logItem.ID, &logItem.EventType, &logItem.SourceIP, &logItem.TargetEndpoint, &logItem.Severity, &logItem.Description, &logItem.CreatedAt)
		if err != nil {
			log.Println("อ่านข้อมูลแถวนี้พลาด:", err)
			continue
		}
		logs = append(logs, logItem)
	}

	// 5. ส่งข้อมูลกลับเป็น JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

// ฟังก์ชัน API รับข้อมูล POST
func createLogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newLog SecurityLog
	err := json.NewDecoder(r.Body).Decode(&newLog)
	if err != nil {
		http.Error(w, "ข้อมูล JSON ไม่ถูกต้อง", http.StatusBadRequest)
		return
	}

	// 1. เตรียมคำสั่ง SQL สำหรับเพิ่มข้อมูล (INSERT)
	// $1, $2... คือตัวแทนที่เราจะเอาข้อมูลมาเสียบ ป้องกันการโดนแฮ็กแบบ SQL Injection
	sqlStatement := `
		INSERT INTO security_logs (event_type, source_ip, target_endpoint, severity, description)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at` // ให้ Postgres ส่ง id และเวลาที่มันสร้างอัตโนมัติกลับมาให้เราด้วย

	var insertedID int
	var createdAt string

	// 2. สั่งรัน SQL ลง Database
	err = db.QueryRow(sqlStatement, newLog.EventType, newLog.SourceIP, newLog.TargetEndpoint, newLog.Severity, newLog.Description).Scan(&insertedID, &createdAt)
	if err != nil {
		log.Println("เกิดข้อผิดพลาดตอนบันทึกลง Database:", err)
		http.Error(w, "เซิร์ฟเวอร์ขัดข้อง ไม่สามารถบันทึกข้อมูลได้", http.StatusInternalServerError)
		return
	}

	// 3. ปริ้นบอกใน Terminal ว่าสำเร็จ
	fmt.Printf("บันทึก Log สำเร็จ! ID: %d | ประเภท: %s\n", insertedID, newLog.EventType)

	// 4. ส่งข้อมูลยืนยันกลับไปหาคนที่เรียก API
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message": "บันทึกข้อมูลลง Database ถาวรเรียบร้อย", "log_id": %d, "timestamp": "%s"}`, insertedID, createdAt)
}

// ฟังก์ชัน API ดึงข้อมูล GET

func main() {
	// 1. ตั้งค่าการเชื่อมต่อ Database
	// *** แก้ไขรหัสผ่านตรงนี้ ***
	connStr := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=security_db sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("ตั้งค่า Database ผิดพลาด: ", err)
	}
	defer db.Close()

	// ทดสอบเคาะประตูก่อนเริ่มทำงาน
	err = db.Ping()
	if err != nil {
		log.Fatal("เชื่อมต่อ Database ไม่สำเร็จ: ", err)
	}
	fmt.Println("🎉 เชื่อมต่อ Database สำเร็จ!")

	// 2. เปิด API รอรับข้อมูล
	http.HandleFunc("/api/v1/logs", createLogHandler)

	// เปิด API เส้นใหม่สำหรับดึงข้อมูล
	http.HandleFunc("/api/v1/logs/all", getLogsHandler)

	fmt.Println("🚀 Security Log API เปิดทำงานแล้วที่ http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("เซิร์ฟเวอร์หยุดทำงาน: ", err)
	}
}
