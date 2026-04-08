package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("--situation1: right data--")
	goodInput := "100"

	number1, err1 := strconv.Atoi(goodInput)

	if err1 != nil {
		fmt.Printf("something wrong: %v\n", err1)
	} else {
		fmt.Printf("Yes that right: %d\n", number1)
	}

	fmt.Println("\n--situation2: wrong data--")
	badInput := "ABC"

	number2, err2 := strconv.Atoi(badInput)

	if err2 != nil {
		fmt.Printf("Nah cant run! cuz %v\n", err2)
	} else {
		fmt.Printf("Yes that right: %d\n", number2)
	}
}
