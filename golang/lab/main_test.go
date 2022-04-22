package main

import "testing"

func BenchmarkDylan(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dylanVersion()
	}
}

func BenchmarkDaochun(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		daochunVersion()
	}
}
