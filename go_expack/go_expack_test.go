package main

import (
	"testing"
	"strings"
)

func TestFizz(t *testing.T) {
	if !strings.EqualFold("FizzBuzz", FizzBuzz(16)) {
		t.Errorf("Correct! Wow, I hope require FizzBuzz")
	}
	if !strings.EqualFold("Fizz", FizzBuzz(4)){
		t.Errorf("Correct! Wow, I hope require Fizz")
	}
	if !strings.EqualFold("Buzz", FizzBuzz(8)) {
		t.Errorf("Correct! Wow, I hope require Buzz")
	}
	if !strings.EqualFold("98", FizzBuzz(100)) {
		t.Errorf("Correct! Wow, I hope require 98")
	}
}
