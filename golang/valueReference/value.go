package main

import "fmt"

type user struct {
	Name string
}

type mm map[string]string

func main() {
	// 结构体
	u := user{
		Name: "daochun",
	}
	fmt.Println("Struct: ")
	changeStructValue(u)
	fmt.Println(u)
	changeStructPointer(&u)
	fmt.Println(u)
	fmt.Println("")

	// Map
	m := mm{"key1": "value1"}
	fmt.Println("Map: ")
	changeMapValue(m)
	fmt.Println(m)
	changeMapPointer(&m)
	fmt.Println(m)
	fmt.Println("")

	// Array
	a := [2]string{"a", "b"}
	fmt.Println("Array")
	changeArrValue(a)
	fmt.Println(a)
	changeArrPointer(&a)
	fmt.Println(a)
	fmt.Println("")
}

func changeStructValue(u user) {
	u.Name = "value changed"
}
func changeStructPointer(u *user) {
	u.Name = "pointer changed"
}

func changeMapValue(m mm) {
	m["key1"] = "change value"
}
func changeMapPointer(m *mm) {
	(*m)["key1"] = "change pointer"
}

func changeArrValue(arr [2]string) {
	arr[1] = "change value"
}
func changeArrPointer(arr *[2]string) {
	arr[1] = "change pointer"
}
