package main

import (
	"testing"

	"go.uber.org/goleak"
)

func TestLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	if err := leak(); err != nil {
		t.Fatal("error not expected")
	}
}

func TestLeak2(t *testing.T) {
	defer goleak.VerifyNone(t)
	if err := leak2(); err != nil {
		t.Fatal("error not expected")
	}
}
