package main

import (
	"encoding/json"
	"fmt"
)

type Employee1 struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `[
		{
			"full_name": "danil syah",
			"email":"danil@mail.com",
			"age":29
		},
		{
			"full_name": "haykal dafiansyah",
			"email":"haykal@mail.com",
			"age":5
		}
	]`

	var result []Employee1

	// decoding JSON Array to Slice of struct
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}
}
