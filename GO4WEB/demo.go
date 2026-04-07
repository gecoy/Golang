//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "GGG", "0800000000", "tanawat2302@gmail.com"})
	fmt.Println(string(data))
}
