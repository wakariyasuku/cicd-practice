package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(20)
	if result != "even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}
