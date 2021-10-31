// Package chapter5 - explains Functions chapter of book with examples
package chapter5

import (
	"fmt"
	"sort"
)

// MyNamedParams - named parameters with struct
type MyNamedParams struct {
	NamedParam1 string
	NamedParam2 int
}

// ReceiveParams - simulate to receive named params with struct type
func ReceiveParams(opts MyNamedParams) {
	// do some stuff on opts variable
	fmt.Println(opts.NamedParam2)
}

// SumUp - is example of variadic params (values)
func SumUp(base int, values ...int) int {
	for _, v := range values {
		base += v
	}
	return base
}

// TestNamedReturn - go support naming return types. This example shows this.
func TestNamedReturn(val1, val2 int) (sum int, multiply int) {
	sum = val1 + val2
	multiply = val1 * val2
	return sum, multiply
}

// TestNamedWithBlankReturn example of returning named params with blank return keyword
func TestNamedWithBlankReturn(val1, val2 int) (sum int, multiply int) {
	sum = val1 + val2
	multiply = val1 * val2
	// this is blank return, and it's very valid in golang
	return
}

// Functions are values. The type of function is built out of the keyword func and the types of the parameters and return values.
func add(val1, val2 int) int {
	return val1 + val2
}

func sub(val1, val2 int) int {
	return val1 - val2
}

func mul(val1, val2 int) int {
	return val1 * val2
}

func div(val1, val2 int) int {
	return val1 / val2
}

// In here, there is map key type is string and value type is function. This proves that function is value
var operationMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

// Expression use named param
type Expression struct {
	FirstOperand  int
	Operation     string
	SecondOperand int
}

// Calculate - is function that takes Expression named param and calculate according operation from operationMap
func Calculate(exp Expression) int {
	function, ok := operationMap[exp.Operation]
	if !ok {
		panic("wrong operation type")
	}
	return function(exp.FirstOperand, exp.SecondOperand)
}

// Function Type declaration
// operationFunction is function type
type operationFunction func(int, int) int

// we declared map with string and operationFunction type
var _ = map[string]operationFunction{
	"+": add,
	"-": sub,
}

// Anonymous Functions

// TestAnonymousFunction - example of anonymous functions in golang
func TestAnonymousFunction() {
	for i := 0; i < 5; i++ {
		func(value int) {
			fmt.Println("Value printed from anonymous function : ", value)
		}(i)
	}
}

// Closure is functions declared inside of functions are able to access and modify variables declared in the outer function.
// Example of closure is Slice function in sort package.

type Person struct {
	Name     string
	Lastname string
}

func TestClosure(people []Person) {
	sort.Slice(people, func(i, j int) bool {
		is := people[i].Name < people[j].Name
		is = people[i].Lastname < people[j].Lastname
		return is
	})
	fmt.Println(people)
}
