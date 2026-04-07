//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	e := employee{}
	// ✅ แก้เป็นแบบนี้
	err := json.Unmarshal([]byte(`{"ID":101,"EmployeeName":"GGG","Tel":"0800000000","Email":"tanawat2302@gmail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", e)
}
