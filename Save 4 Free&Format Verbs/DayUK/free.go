package main

import (
	"bufio"
	"fmt"
	"os"

	// "strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// func getInput(promt string) float64 {
// 	fmt.Printf("%v", promt)
// 	input, _ := reader.ReadString('\n')
// 	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
// 	if err != nil {
// 		message, _ := fmt.Scanf("%v must number only", promt)
// 		panic(message)
// 	}
// 	return value
// }

func main() {
	fmt.Println("Enter your Grade: ")
	input, _ := reader.ReadString('\n')
	myGrade := strings.TrimSpace(input)
	myGrade = strings.ToUpper(myGrade)
	switch myGrade {
	case "A":
		fmt.Println("Cool!")
	case "B":
		fmt.Println("Good Job!")
	case "C":
		fmt.Println("Well")

	default:
		fmt.Println("Nah")
	}
}
