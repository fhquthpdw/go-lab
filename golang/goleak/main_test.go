package main

import (
	"testing"

	"go.uber.org/goleak"
)

func TestLeakWithGoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	leak()
}
