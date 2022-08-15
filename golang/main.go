package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	CATDog string `json:"catDog"`
}

func main() {
	// map to struct
	var result Person
	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"catDog": "cat dog",
	}
	if err := mapstructure.Decode(input, &result); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("first round:")
		fmt.Println(result.Name)
		fmt.Println(result.Age)
		fmt.Println(result.CATDog)
	}

	// struct to map
	result2 := make(map[string]interface{})
	input2 := Person{
		Name:   "name2",
		Age:    182,
		CATDog: "dog cat 2",
	}

	if err := mapstructure.Decode(input2, &result2); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("")
		fmt.Println("second round:")
		fmt.Println(result2["Name"])
		fmt.Println(result2["Age"])
		fmt.Println(result2["CATDog"])
	}

}
