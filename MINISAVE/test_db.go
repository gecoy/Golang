//go:build ignore

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // ต้องใส่ _ ไว้ข้างหน้าเพื่อบอก Go ว่าเราโหลด driver มาเฉยๆ แต่ยังไม่ได้เรียกใช้ฟังก์ชันมันตรงๆ
)

func main() {
	// 1. กำหนดข้อมูลการเชื่อมต่อ (เปลี่ยน password เป็นของคุณเอง)
	// รูปแบบ: host=... port=... user=... password=... dbname=... sslmode=disable
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=security_db sslmode=disable"

	// 2. สั่งเปิดประตู (แต่ยังไม่ได้เดินเข้าไป)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("เกิดข้อผิดพลาดในการเปิดฐานข้อมูล: ", err)
	}
	defer db.Close() // สั่งว่าทำงานเสร็จอย่าลืมปิดประตูด้วย

	// 3. ลองเคาะประตูจริงๆ (Ping) ว่าเชื่อมต่อได้ไหม
	err = db.Ping()
	if err != nil {
		log.Fatal("เชื่อมต่อ Database ไม่สำเร็จ! ตรวจสอบรหัสผ่านหรือชื่อฐานข้อมูล: ", err)
	}

	fmt.Println("🎉 เชื่อมต่อ PostgreSQL สำเร็จแล้ว!!")
}
