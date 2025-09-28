package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	// arrange
	s := "(21+33)*3.2-4+(5*2-6)*2="
	exp := 176.8

	// act
	act, err := Calc(s)

	// assert
	if act != exp {
		t.Errorf("Calc(%s) = %f; expected: %d", s, act, exp)
	} else if err != nil {
		t.Error(err)
	}

}
