package main

import "fmt"

type Status int

const (
	Nothing = iota
	Fizz
	Buzz
	FizzBuzz
)

func CheckFizzBuzz(i int) int {
	if i%3 == 0 && i%5 == 0 {
		return FizzBuzz
	}
	if i%3 == 0 {
		return Fizz
	}
	if i%5 == 0 {
		return Buzz
	}
	return Nothing
}

func main() {
	fmt.Println(CheckFizzBuzz(5))
}
