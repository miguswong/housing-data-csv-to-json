package main

import (
	"slices"
	"testing"
)

// Tests to make sure that fileExists function is working
func TestFileExists(t *testing.T) {
	fileName := "housesInput.csv"
	exp := true
	res := fileExists(fileName)

	if res != exp {
		t.Errorf("Expected %t, got %t instead.", exp, res)
	}
}

func TestCreateHousingList(t *testing.T) {
	testData := [][]string{
		{"value", "income", "age", "rooms", "bedrooms", "pop", "hh", "Extra Column"},
		{"1", "1", "1", "1", "1", "1", "1", "18"},
		{"2", "2", "2", "2", "2", "2", "2"},
	}

	exp := []housingRecord{
		{Value: 1, Income: 1, Age: 1, Rooms: 1, Bedrooms: 1, Pop: 1, Hh: 1},
		{Value: 2, Income: 2, Age: 2, Rooms: 2, Bedrooms: 2, Pop: 2, Hh: 2},
	}

	res := createHousingList(testData)
	if !slices.Equal(exp, res) {
		t.Errorf("Expected:\n %v\n Instead, this was the result:\n%v", exp, res)
	}

}
