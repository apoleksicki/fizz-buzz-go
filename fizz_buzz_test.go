package main

import (
	"fmt"
	"testing"
)

func TestCheckFizzBuzz(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{1, Nothing},
		{2, Nothing},
		{7, Nothing},
		{11, Nothing},
		{3, Fizz},
		{12, Fizz},
		{27, Fizz},
		{5, Buzz},
		{10, Buzz},
		{20, Buzz},
		{25, Buzz},
		{0, FizzBuzz},
		{15, FizzBuzz},
		{60, FizzBuzz},
		{120, FizzBuzz},
		{150, FizzBuzz},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := CheckFizzBuzz(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
