package main

import "fmt"

func main() {
	secretPin := 1234
	maxAttempts := 3 //โอกาส 3 ครั้ง
	isLocked := true //ประตูยังล็อคอยู่

	fmt.Println("--- Security System: Restricted Area ---")

	//Loop i เริ่มที่ 1; ทำไปจนถึงรอบที่ 3; เพิ่มรอบทีละ 1
	for i := 1; i <= maxAttempts; i++ {
		fmt.Printf("Attemt #%d Enter PIN: ", i)

		var inputPin int
		fmt.Scan(&inputPin)

		if inputPin == secretPin {
			fmt.Println("Access Granted! Welcome, Admin.")
			isLocked = false //unlocked success
			break            //คำสั่งสำคัญ: เจอ Break จะกระโดดออกจาก Loop ทันที (ไม่ต้องทำรอบที่เหลือ)
		} else {
			fmt.Println("Access Denied! Wrong PIN.")
		}
	}

	// เช็คผลลัพธ์หลังจากจบ Loop (ไม่ว่าจะถูก หรือ ผิดจนครบ)
	if isLocked {
		fmt.Println("\n SYSTEM LOCKED! Too many failed attempts.")
		fmt.Println("Police are on the way...")
	}
}
