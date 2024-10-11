package main

import "testing"

func TestFileExsits(t *testing.T) {
	fileName := "housesInput.csv"
	exp := true
	res := fileExists(fileName)

	if res != exp {
		t.Errorf("Expected %t, got %t instead.", exp, res)
	}
}
