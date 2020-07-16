package main_test

import "testing"

func TestAdd(t *testing.T){
	got:= 2 + 2
	expected:= 4

	if got != expected {
		t.Errorf("Did Not get what expected. Got: %v, expected: %v", got, expected)
	}
}
