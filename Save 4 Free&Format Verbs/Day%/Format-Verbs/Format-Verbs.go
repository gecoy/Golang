package main

import "fmt"

func main() {
	name := "Somchai"
	age := 25
	score := 95.50

	// ใช้ %v (ครอบจักรวาล)
	fmt.Printf("Name: %v, Age: %v\n", name, age)
	// ผลลัพธ์: Name: Somchai, Age: 25

	// ใช้ให้ตรงประเภท (%s, %d, %.2f)
	fmt.Printf("Name: %s, Age: %d, Score: %.2f\n", name, age, score)
	// ผลลัพธ์: Name: Somchai, Age: 25, Score: 95.50

	// ใช้ %T เพื่อดูชนิดตัวแปร
	fmt.Printf("type of name is: %T\n", name)
	fmt.Printf("type of score is: %T\n", score)
	// ผลลัพธ์:
	// type of name is: string
	// type of score is: float64
}
