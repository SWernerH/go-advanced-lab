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

func MakeAccumulator(initial int) (func(int), func(int), func() int) {

	value := initial
	add := func(x int) {
		value += x
	}
	subtract := func(x int) {
		value -= x
	}
	get := func() int {
		return value
	}
	return add, subtract, get
}

//part 3

func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = operation(num)
	}
	return result
}

func Filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range nums {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

func Reduce(nums []int, initial int, operation func(int, int) int) int {
	acc := initial
	for _, num := range nums {
		acc = operation(acc, num)
	}
	return acc
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

//part 4

func ExploreProcess(){

	fmt.Println("Process info:")
	fmt.Println("current process ID:", os.Getpid())
	fmt.Println("parent process ID:", os.Getppid())

	data := []int{1,2,3,4,5}
	fmt.Printf("Memory address of data slice: %p\n", &data)
	fmt.Printf("Memory address of first element: %p\n", &data[0])

	fmt.Println("no other processes canaccess this memory adress")
}


//part 5

func DoubleValue(x int) {
	x = x * 2
}

func DoublePointer(x *int) {
	*x = *x * 2
}

func CreateOnStack() int {
	x := 42
	return x
}

func CreateOnHeap() *int{
	x := 20
	return &x
}

func SwapValues(a, b int) (int, int) {
	return b, a
}

func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}

func AnalyzeEscape() {
	CreateOnStack()
	CreateOnHeap()
}

//part 6

func main(){
	ExploreProcess()

	fmt.Println("\n========Math operations========")
	f0, _ := Factorial(0)
	f5, _ := Factorial(5)
	f10, _ := Factorial(10)

	fmt.Println("Factorial(0):", f0)
	fmt.Println("Factorial(5):", f5)
	fmt.Println("Factorial(10):", f10)

	p17, _ := IsPrime(17)
	p18, _ := IsPrime(18)
	p20, _ := IsPrime(20)

	fmt.Println("IsPrime(17):", p17)
	fmt.Println("IsPrime(18):", p18)
	fmt.Println("IsPrime(20):", p20)

	power1, _ := Power(2, 3)
	power2, _ := Power(5, 4)

	fmt.Println("Power(2,3):", power1)
	fmt.Println("Power(5,4):", power2)


	fmt.Println("\n********Counter Demonstration********")

	counter1 := MakeCounter(0)
	counter2 := MakeCounter(100)
	
	fmt.Println("Counter1:", counter1())
	fmt.Println("Counter2:", counter2())
	fmt.Println("Counter1:", counter1())
	fmt.Println("Counter2:", counter2())

	doubler := MakerMuntiplier(14)
	tripler := MakerMuntiplier(3)
	
	fmt.Println("Doubler(5):", doubler(5))
	fmt.Println("Tripler(5):", tripler(5))

	fmt.Println("\n--------Higer-order functions --------")

	nums := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("Original numbers:", nums)

	squared := Apply(nums, func(x int) int {return x * x})
	fmt.Println("Squared numbers:", squared)

	evens := Filter(nums, func(x int) bool {return x%2 == 0})
	fmt.Println("Even numbers:", evens)

	sum := Reduce(nums, 0, func(acc, curr int) int {return acc + curr})
	fmt.Println("Sum of numbers:", sum)

	DuobleThenAddTen := Compose(func(x int)int {return x * 10}, func(x int) int {return x + 2})
	fmt.Println("Double then add ten to 5:", DuobleThenAddTen(5))

	fmt.Println("\n~~~~~~~~Pointer Demonstration~~~~~~~~")

	a := 5
	b := 10

	fmt.Println("Before SwapValues: a =", a, ", b =", b)
	a, b = SwapValues(a, b)
	fmt.Println("After SwapValues: a =", a, ", b =", b)

	c := 15
	d := 20

	fmt.Println("Before SwapPointers: c =", c, ", d =", d)
	SwapPointers(&c, &d)
	fmt.Println("After SwapPointers: c =", c, ", d =", d)

	AnalyzeEscape()
}