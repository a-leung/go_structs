package main

import "testing"

func TestDice(t *testing.T) {
	dice := &Dice{}

	dice.SetSeed(1)

	if dice.GetSeed() != 1 {
		t.Error("Issue with set seed")
	}

	dice.SetSeed(2)

	if dice.GetSeed() == 1 {
		t.Error("Issue with set seed")
	}
}
