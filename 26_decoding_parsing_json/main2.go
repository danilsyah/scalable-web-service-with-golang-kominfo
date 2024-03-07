package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `
		{
			"full_name": "Haykal Ganteng",
			"email": "haykal@mail",
			"age": 5
		}
	`

	var result map[string]interface{}

	// decoding JSON to MAP
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name:", result["full_name"])
	fmt.Println("email:", result["email"])
	fmt.Println("age:", result["age"])
}
