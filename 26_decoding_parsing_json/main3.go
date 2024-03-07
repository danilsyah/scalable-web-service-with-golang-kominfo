package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `
		{
			"full_name": "Nufika fitriani",
			"email": "fika@mail.com",
			"age": 29
		}
	`

	var temp interface{}

	// decoding JSON to Empty Interface
	var err = json.Unmarshal([]byte(jsonString), &temp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result = temp.(map[string]interface{})

	fmt.Println("full_name:", result["full_name"])
	fmt.Println("email:", result["email"])
	fmt.Println("age:", result["age"])
}
