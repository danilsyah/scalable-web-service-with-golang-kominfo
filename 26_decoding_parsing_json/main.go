package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `
		{
			"full_name": "Danil Syah",
			"email": "danil@mail.co.id",
			"age": 29
		}
	`

	var result Employee

	// decoding JSON to Struct
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name:", result.Fullname)
	fmt.Println("email:", result.Email)
	fmt.Println("age:", result.Age)
}
