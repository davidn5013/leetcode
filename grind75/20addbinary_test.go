package grind75

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	a, b := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	got := AddBinary(a, b)
	want := "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"
	if got != want {
		t.Errorf("Addbinary want \n[%s]\n;got \n[%s]\n",
			want, got)
	}
}

func TestAddBinary2(t *testing.T) {

	a, b := "10101010101010101010", "101"
	got := AddBinary(a, b)
	want := "1111"
	if got != want {
		t.Errorf("Addbinary want \n[%s]\n;got \n[%s]\n",
			want, got)
	}
}

func TestBinAddInString(t *testing.T) {
	got, cr := binAddInString('0', '0', '0')
	want := '0'
	want2 := '0'

	if got != want || cr != want2 {
		t.Errorf("got %c cr %c ;wanted %c %c\n", got, cr, want, want2)
	}

	got, cr = binAddInString('0', '1', '0')
	want = '1'
	want2 = '0'

	if got != want || cr != want2 {
		t.Errorf("got %c cr %c ;wanted %c %c\n", got, cr, want, want2)
	}

	got, cr = binAddInString('0', '1', '1')
	want = '0'
	want2 = '1'

	if got != want || cr != want2 {
		t.Errorf("got %c cr %c ;wanted %c %c\n", got, cr, want, want2)
	}

	got, cr = binAddInString('1', '1', '1')
	want = '1'
	want2 = '1'

	if got != want || cr != want2 {
		t.Errorf("got %c cr %c ;wanted %c %c\n", got, cr, want, want2)
	}
}

// Just for fun are what is faster |1==0 or %2==0?

func BenchmarkBitEven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := 0
		if 10&1 == 0 {
			c++
		}
	}
}

func BenchmarkModEven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := 0
		if 10%2 == 0 {
			c++
		}
	}
}
