package main

import (
	"testing"
)

//part 1 tests
func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{"Factorial of 5", 5, 120, false},
		{"Factorial of 0", 0, 1, false},
		{"Factorial of negative", -3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{"IsPrime of 7", 7, true, false},
		{"IsPrime of 4", 4, false, false},
		{"IsPrime of 1", 1, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool

	}{
		{"Power 2^3", 2, 3, 8, false},
		{"Power 5^0", 5, 0, 1, false},
		{"Power with negative exponent", 2, -2, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Power() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}

//part 2 tests

func TestMakeCounter(t *testing.T) {
	counter1 := MakeCounter(0)
	counter2 := MakeCounter(10)

	if counter1() != 1 {
		t.Errorf("expected 1, got %d", counter1())
	}
	if counter1() != 2 {
		t.Errorf("expected 2, got %d", counter1())
	}
	if counter2() != 11 {
		t.Errorf("expected 11, got %d", counter2())
	}
	if counter2() != 3 {
		t.Errorf("expected 3, got %d", counter2())
	}
}

func TestMakeMultiplier(t *testing.T){
	tests := []struct {
		name     string
		factor   int
		inputs   int
		want	 int
	}{
		{"Multiply by 2", 2, 5, 10},
		{"Multiply by 3", 3, 4, 12},
		{"Multiply by 0", 0, 10, 0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiplier := MakerMuntiplier(tt.factor)
			got := multiplier(tt.inputs)
			if got != tt.want {
				t.Errorf("MakerMultiplier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	add, subtract, get := MakeAccumulator(10)
	
	add(50)
	if got := get(); got != 60 {
		t.Errorf("after adding 50, expected 60, got %d", got)
	}
	
	subtract(20)
	if got := get(); got != 40 {
		t.Errorf("after subtracting 20, expected 40, got %d", got)
	}
}
//part 3 tests

func TestApply(t *testing.T) {
	test := []struct {
		name     string
		input   []int
		operation func(int) int
		want 	[]int
	}{
		{"Double values", []int{1, 2, 3}, func(x int) int { return x * 2 }, []int{2, 4, 6}},
		{"Square values", []int{1, 2, 3}, func(x int) int { return x * x }, []int{1, 4, 9}},
		{"negate values", []int{1, -2, 3}, func(x int) int { return -x }, []int{-1, 2, -3}},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.input, tt.operation)
			if len (got) != len(tt.want) {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Apply() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	test := []struct {
		name     string
		input   []int
		predicate func(int) bool
		want 	[]int
	}{
		{"Even numbers", []int{1, 2, 3, 4, 5}, func(x int) bool { return x%2 == 0 }, []int{2, 4}},
		{"Greater than 3", []int{1, 2, 3, 4, 5}, func(x int) bool { return x > 3 }, []int{4, 5}},
		{"grater than 10", []int{5, 10, 15, 20}, func(x int) bool { return x > 10 }, []int{15, 20}},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.input, tt.predicate)
			if len (got) != len(tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Filter() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	test := []struct {
		name		string
		input		[]int
		initial		int
		operation	func(int, int) int
		want		int
	}{
		{"Sum values", []int{1, 2, 3, 4}, 0, func(acc, curr int) int { return acc + curr }, 10},
		{"Product values", []int{1, 2, 3, 4}, 1, func(acc, curr int) int { return acc * curr }, 24},
		{"Max value", []int{1, 5, 3, 4}, 0, func(acc, curr int) int {
			if curr > acc {
				return curr
			}
			return acc
		}, 5},
		{"Min", []int{5,3, 2, 4}, 999, func(acc, curr int) int {
			if curr < acc {
				return curr
			}
			return acc
		}, 2},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.input, tt.initial, tt.operation)
			if got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCompose(t *testing.T) {
	test := []struct {
		name string
		f    func(int) int
		g    func(int) int
		input int
		want  int
	}{
		{"Duoble then add 2", func(x int) int { return x * 2 }, func(x int) int { return x + 2 }, 5, 12},
		{"Squere then add one", func(x int) int { return x * x }, func(x int) int { return x + 1 }, 3, 10},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			composed := Compose(tt.f, tt.g)
			got := composed(tt.input)
			if got != tt.want {
				t.Errorf("Compose() = %v, want %v", got, tt.want)
			}
		})
	}
}

//part 5 tests

func TestSwapValues(t *testing.T) {
	a, b := 5, 10
	x, y := SwapValues(a, b)
	if x != 10 || y != 5 {
		t.Errorf("SwapValues() = (%v, %v), want (10, 5)", x, y)
	}
	if a != 5 || b != 10 {
		t.Errorf("original values changed: a = %v, b = %v", a, b)
	}
}

func TestSwapPointers(t *testing.T)	{
	a, b := 5, 10
	SwapPointers(&a, &b)
	if a != 10 || b != 5 {
		t.Errorf("SwapPointers() = %d, %d; want (10, 5)", a, b)
	}
}