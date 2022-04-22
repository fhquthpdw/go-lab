package main

import "testing"

func BenchmarkSliceAppend(b *testing.B) {
	var a []int
	for i := 0; i < b.N; i++ {
		a = append(a, i)
	}
}

func BenchmarkSliceAppend2(b *testing.B) {
	a := make([]int, 0, 0)
	for i := 0; i < b.N; i++ {
		a = append(a, i)
	}
}

/*
func BenchmarkSliceSet(b *testing.B) {
	var a []int
	for i := 0; i < b.N; i++ {
		a[i] = i
	}
}

type b struct {
	name    string
	age     string
	address string
}

var a = []b{
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
	{name: "fhquthpdw", age: "ee", address: "address"},
}

//
func BenchmarkFor(b *testing.B) {
	l := len(a)
	for i := 0; i < b.N; i++ {
		for idx := 0; idx < l; idx++ {
			v := a[idx]
			showV(v)
		}
	}
}

func BenchmarkRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range a {
			showV(v)
		}
	}
}

func BenchmarkRange2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for aa := range a {
			showV(a[aa])
		}
	}
}

func showV(_ interface{}) {

}
*/
