package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid data", 10, 2, 5, false},
	{"valid data", 10, 0, 0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)

		if err != nil {
			if !tt.isErr {
				t.Error("Gpt an err")
			}
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)

	if err != nil {
		t.Error("Got an error when it should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)

	if err == nil {
		t.Error("Did not get an error when it should have")
	}
}
