package main

import "testing"

func TestCalculateChange(t *testing.T) {
	expected := make(map[coin]int)
	expected[coins[0]] = 1
	expected[coins[1]] = 1


	res := calculateChange(1.50)
	if res[coins[0]] != 1  && res[coins[1]] != 1{
		t.Error("Error on changing")
	}
}