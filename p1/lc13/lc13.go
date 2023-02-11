// package lc13 LeetCode 0013 Roman to Integer
// Latin number to integer
package lc13

import "log"

// RomanToInt convert roman number to int
func RomanToInt(s string) int {
	r := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	finddigit := func(s string) int {
		if v, ok := r[s]; ok {
			return v
		}
		return 0
	}

	var result int

	for i := 0; i < len(s); i++ {
		romChar := string(s[i])

		var romNextChar string
		if i < len(s)-1 {
			romNextChar = string(s[i+1])
		}

		two := romChar + romNextChar
		addTwo := finddigit(two)

		if addTwo != 0 {
			log.Println(two, addTwo)
			result += addTwo
			i++
			continue
		}

		add := finddigit(romChar)
		log.Println(romChar, add)
		result += add
	}
	return result
}
