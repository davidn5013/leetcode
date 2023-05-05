package p1

import "strconv"

// FizzBuzz 3=Fizz 5=Buzz 15=FizzBuzz
/*
Given an integer n, return a string array answer (1-indexed) where:
answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true. */
func FizzBuzz(n int) (result []string) {
	var s string
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0:
			s += "Fizz"
		case i%5 == 0:
			s += "Buzz"
		default:
			s += strconv.Itoa(i)
		}
		result = append(result, s)
		s = ""
	}
	return result
}
