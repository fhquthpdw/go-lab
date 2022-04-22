package main

import (
	"net/http"
	"testing"
)

func BenchmarkHttprouter(b *testing.B){
	b.ResetTimer()
	for i:=0; i<b.N; i++{
		http.Get("http://localhost:19292/menu/get/daochun")
	}
}

func BenchmarkEcho(b *testing.B){
	b.ResetTimer()
	for i:=0; i<b.N; i++{
		http.Get("http://localhost:19191/menu/get/daochun")
	}
}

func BenchmarkGin(b *testing.B){
	b.ResetTimer()
	for i:=0; i<b.N; i++{
		http.Get("http://localhost:19393/menu/get/daochun")
	}
}

func BenchmarkIris(b *testing.B){
	b.ResetTimer()
	for i:=0; i<b.N; i++{
		http.Get("http://localhost:19494/menu/get/daochun")
	}
}

func BenchmarkMartini(b *testing.B){
	b.ResetTimer()
	for i:=0; i<b.N; i++{
		http.Get("http://localhost:19595/menu/get/daochun")
	}
}

/*
func BenchmarkBeego(b *testing.B){
	b.ResetTimer()
	for i:=0; i<b.N; i++{
		http.Get("http://localhost:19696/menu/get/daochun")
	}
}
*/