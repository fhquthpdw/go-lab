package main

import (
	"testing"

	"go.uber.org/goleak"
)

// 可以通过 go test -run Test_queryAll 检测出来
func Test_queryAll(t *testing.T) {
	defer goleak.VerifyNone(t)

	tests := []struct {
		name string
	}{
		{"first"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryAll()
		})
	}
}

// 测试不出来
func TestWaitingGoroutineLeak(t *testing.T) {
	defer goleak.VerifyNone(t)

	tests := []struct {
		name string
	}{
		{"leak test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WaitingGoroutineLeak()
		})
	}
}
