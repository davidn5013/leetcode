package p1

import (
	"fmt"
	"math"
)

// usqrtGo idiomatic go
func usqrtGo(val uint) uint {
	return uint(math.Sqrt(float64(val)))
}

// usqrt div 5 time method
func usqrt4(val uint) uint {
	if val < 2 {
		return val
	}

	var a, b uint

	a = 1255 // starting point is relatively unimportant

	for i := 1; i <= 20; i++ {
		b = val / a
		a = (a + b) / 2
	}

	return a
}

func main() {
	fmt.Printf("%d %d\n", usqrtGo(99999999999), usqrt4(99999999999))
}

/*

	// static will allow inlining
	static unsigned usqrt4(unsigned val) {
		unsigned a, b;

		if (val < 2) return val; // avoid div/0

		a = 1255;       // starting point is relatively unimportant

		b = val / a; a = (a+b) /2;
		b = val / a; a = (a+b) /2;
		b = val / a; a = (a+b) /2;
		b = val / a; a = (a+b) /2;

		return a;
	}

	unsigned char bit_width(unsigned long long x) {
		return x == 0 ? 1 : 64 - __builtin_clzll(x);
	}

	// implementation for all unsigned integer types
	unsigned sqrt(const unsigned n) {
		unsigned char shift = bit_width(n);
		shift += shift & 1; // round up to next multiple of 2

		unsigned result = 0;

		do {
			shift -= 2;
			result <<= 1; // leftshift the result to make the next guess
			result |= 1;  // guess that the next bit is 1
			result ^= result * result > (n >> shift); // revert if guess too high
		} while (shift != 0);

		return result;
	}

	static uint32_t usqrt4_6(uint32_t val) {
		uint32_t a, b;

		if (val < 2) return val; // avoid div/0
		a = 1255;       // starting point is relatively unimportant
		b = val / a; a = (a + b)>>1;
		b = val / a; a = (a + b)>>1;
		b = val / a; a = (a + b)>>1;
		b = val / a; a = (a + b)>>1;
		if (val < 20000) {
			b = val / a; a = (a + b)>>1;    // < 17% error Max
			b = val / a; a = (a + b)>>1;    // < 5%  error Max
		}
		return a;
	}

*/
