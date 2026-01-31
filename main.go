package main

import(
	"fmt"
	"errors"
	"os"
)

//part 1

func Factorial(n int) (int, error) {

	if n < 0 {
		return 0, errors.New("factorial not defined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	result := 1
	for i := 1; i <=n ; i ++ {
		result *= i
	}

	return result, nil
}

func IsPrime(n int) (bool, error) {

	if n < 2 {
		return false, errors.New("prime check requires n >= 2")
	}
	if n == 2 {
		return true, nil
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0{
			return false, nil
		}
	}
	return true, nil
}

func Power (base, exponent int)(int, error){

	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base 
	} 
	return result, nil
}


//part 2

func MakeCounter(start int) func() int {

	counter := start
	return func() int {
		counter++
		return counter
	}
}

func MakerMuntiplier(factor int) func(int) int {

	return func(x int) int{
		return x * factor
	}
}

func MakeAccumulator(initial int) func(int) func(int) int func(int) int {

	value := initial
	add := funvc(x int) {
		value += x
	}
	subtract func(x int) {
		value -= x
	}
	get := func() int {
	rerturn value
	}
	return add, subtract, get
}

//part 3

func Apply(nums[]int, operations func(int)int) []int {

	result := make([]int, len(nums))
	for i, num = range nums {
		result[i] = operations(num)
	}
	return result
}

func filter(nums []int, predicate func (int) bool) []int {

	var result []int
	for _, nums := range nums {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

func Resuce(nums []int, initial in, operations func(int, int)int) int {

	acc := initial
	for _, nums := range nums {
		acc = operations(acc, num)
	}
	return acc
}

func compose (f func(int)int, g func(int)int,) func(int) int {

	return func(x int) int {
		return f(g(x))
	}
}

