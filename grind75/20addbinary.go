package grind75

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// AddBinary add too a string of binaries "11"+"1"="100"
// Arbitrary length
func AddBinary(a string, b string) (res string) {

	// TODO "00001010" , "101" will not work
	// TODO "10101000", "101" giv "111110000" instead of "10101101"

	// read string from right to left
	// letter by letter add 1+0=1 0+0=0 1+1=1 sig=1
	// print string
	//
	// How read a string from right to left

	// I want the shogt in the inner loop
	if len(b) > len(a) {
		temp := a
		a = b
		b = temp
	}

	var sb strings.Builder
	sb.Grow(200)

	// carrier
	cr := '0'
	var t rune
	aMax, bMax := len(a)-1, len(b)-1
	for i1, i2 := aMax, bMax; i1 >= 0; i1-- {

		// Add with carrier
		if i2 >= 0 {
			as := rune(a[i1])
			bs := rune(b[i2])
			t, cr = binAddInString(as, bs, cr)
			sb.WriteByte(byte(t))
		} else {
			// Carrier the carrier out
			t, cr = binAddInString(rune(a[i1]), '0', cr)
			sb.WriteByte(byte(t))
		}
		if i2 >= 0 {
			i2--
		}
	}
	return sb.String()
}

func binAddInString(a, b, cr rune) (res, crOut rune) {
	sum := 0
	res = '0'
	crOut = '0'

	for _, v := range []rune{a, b, cr} {
		i, err := strconv.ParseInt(string(v), 2, 64)
		if err != nil {
			log.Fatal(v, err)
		}
		sum += int(i)
	}

	binStr := fmt.Sprintf("%b", sum)

	lenbinStr := len(binStr)
	if lenbinStr == 1 {
		res = rune(binStr[0])
	} else if lenbinStr == 2 {
		res = rune(binStr[1])
		crOut = rune(binStr[0])
	} else {
		panic("Wrong binary number")
	}

	return res, crOut
}

/*
func binAddInStringDefekt(a, b, cr rune) (res, crOut rune) {
	res, crOut = '0', '0' // do not return runes with zero

	// don't allow zero input carrier
	// I do note want  as cr == 0
	if cr != '0' && cr != '1' {
		cr = '0'
	}

	if a == '0' && b == '0' && cr == '0' {
		res = '0'
		crOut = '0'
		return res, crOut
	}

	if a == '1' && b == '0' && cr == '0' {
		res = '1'
		crOut = '0'
		return res, crOut
	}

	if a == '0' && b == '1' && cr == '0' {
		res = '1'
		crOut = '0'
		return res, crOut
	}

	if a == '1' && b == '1' && cr == '0' {
		res = '0'
		crOut = '1'
		return res, crOut
	}

	if a == '0' && b == '0' && cr == '1' {
		res = '1'
		crOut = '0'
		return res, crOut
	}

	if a == '0' && b == '1' && cr == '1' {
		res = '0'
		crOut = '1'
		return res, crOut
	}

	if a == '1' && b == '0' && cr == '1' {
		res = '0'
		crOut = '1'
		return res, crOut
	}

	if a == '1' && b == '1' && cr == '1' {
		res = '1'
		crOut = '1'
		return res, crOut
	}

	return res, crOut
}*/
