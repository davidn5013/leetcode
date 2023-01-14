package grind75

import (
	"errors"
	"log"
	"strconv"
)

// leetcode 150. Evaluate Reverse Polish Notation
// Grind75 number 33

// Eval reverse polish notation like
// "2","1","+","3","*" to 6
func EvalRPN(tokens []string) (res int) {
	var values []int

	for len(tokens) > 0 {

		v := tokens[0]
		tokens = tokens[1:]

		switch {
		case v == "+", v == "-", v == "*", v == "/":
			values = popOp(values, v)
		default:
			i, err := strconv.Atoi(v)
			if err != nil {
				err := errors.New("token is not a number and not operator")
				log.Println(err)
			}
			values = append(values, i)
		}
		res = values[0]
	}

	return res
}

func popOp(values []int, op string) []int {
	y, values := pop(values)
	x, values := pop(values)
	switch op {
	case "+":
		values = append(values, x+y)
	case "-":
		values = append(values, x-y)
	case "/":
		values = append(values, x/y)
	case "*":
		values = append(values, x*y)
	}
	return values
}

func pop(values []int) (x int, resValues []int) {
	if len(values) < 1 {
		err := errors.New("Less then two values on value stack when try to do op")
		log.Println(err)
		return 0, []int{0}
	}

	x = values[len(values)-1]
	resValues = values[:len(values)-1]
	return x, resValues
}

/* Fastes and best look solution on leetcode
 */

/*
func evalRPN(tokens []string) int {
    operators := map[string]func(int,int)int{
        "+": sum,
        "-": substract,
        "*": multiply,
        "/": div,
    }

    nums := []int{}
    for _, v := range tokens {
        if _, ok := operators[v]; ok {
            a := pop(&nums)
            b := pop(&nums)
            res := operators[v](b, a)
            nums = append(nums, res)
        }else{
            n, _ := strconv.Atoi(v)
            nums = append(nums, n)
        }
    }

    return nums[0]
}

func pop(nums *[]int) int{
    pop := (*nums)[len(*nums)-1]
    *nums = (*nums)[:len(*nums)-1]
    return pop
}

func multiply(a, b int) int {
    return a * b
}

func sum(a, b int) int {
    return a + b
}

func substract(a, b int) int {
    return a - b
}

func div(a, b int) int {
    return a / b
}

*/
