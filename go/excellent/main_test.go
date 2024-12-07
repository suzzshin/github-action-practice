package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)

	if result != "Even" {
		t.Errorf("EvenOrOdd(10) = %s; want Even", result)
	}
}
