package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(15, 15)

	if total != 30 {
		t.Errorf("Sum result is invalid: Result %d. Expected %d", total, 30)
	}
}
