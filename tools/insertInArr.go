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
