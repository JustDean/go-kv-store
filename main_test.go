package main

import "testing"

func TestPut(t *testing.T) {
	name := "Test Put"
	res := Put("0", "Aboba")

	if res != nil {
		t.Fatalf(`Missmatch in function %s. Result is %s. Expected nil`, name, res)
	}
}

// TODO tests
