package main

import (
	"encoding/json"
	"fmt"
)

type S1 struct {
	Name string `json:"name"`
}

func (s S1) PrintName() {
	fmt.Println(s.Name)
}

type S2 struct {
	S1
	//Name []string `json:"name"`
}

func main() {
	s1 := S1{Name: "name of s1"}
	s2 := S2{
		S1: s1,
	}

	s2.PrintName()
	s2.Name = "name of s3"

	s, _ := json.Marshal(s2)

	fmt.Println(string(s))
}
