package tools

import "fmt"

// InsertIntAt insert []int at position in array
func InsertIntAt(inp []int, position int, value []int) (res []int, err error) {
	if position < 0 || position > len(inp) {
		err := fmt.Errorf("InsertAt: position %d out of range %d", position, len(inp))
		return nil, err
	}

	elements := len(value)
	res = append(inp[:position], append(make([]int, elements), inp[position:]...)...)

	for _, v := range value {
		res[position] = v
		position++
		if position > position+elements {
			break
		}
	}
	return res, nil

}

// InsertAt insert []genric at position in array
func InsertAt[K comparable](inp []K, position int, value []K) (res []K, err error) {
	if position < 0 || position > len(inp) {
		err := fmt.Errorf("InsertAt: position %d out of range %d", position, len(inp))
		return nil, err
	}

	elements := len(value)
	res = append(inp[:position], append(make([]K, elements), inp[position:]...)...)

	for _, v := range value {
		res[position] = v
		position++
		if position > position+elements {
			break
		}
	}
	return res, nil

}

// EqualMatrix return two are the same [][]int
func EqualMatrix(a, b [][]int) (res bool) {
	if len(a) != len(b) {
		return false
	}
	res = true

	for idx1, arr1 := range a {
		if len(arr1) != len(b[1]) {
			res = false
			break
		}
		arr2 := b[idx1]
		for idx2, value1 := range arr1 {
			value2 := arr2[idx2]
			if value1 != value2 {
				res = false
				goto endloop
			}
		}
	}
endloop:
	return res
}

// Copy2dInt make copy [][]int 2d array.
func Copy2dInt(src [][]int) (des [][]int) {
	// copy mat to res
	des = make([][]int, len(src))
	for i := range src {
		des[i] = make([]int, len(src[i]))
		copy(des[i], src[i])
	}
	return des
}
