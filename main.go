package main

import (
	"book/chapter5"
	"fmt"
)

func main() {
	chapter5Examples()
}

func chapter5Examples() {
	finalSum := chapter5.SumUp(2, 12, 12, 33, 44)
	fmt.Println(finalSum)
	fmt.Println(chapter5.TestNamedReturn(2, 4))
	fmt.Println(chapter5.TestNamedWithBlankReturn(4, 6))
	fmt.Println(chapter5.Calculate(chapter5.Expression{
		FirstOperand:  3,
		Operation:     "+",
		SecondOperand: 4,
	}))
	chapter5.TestAnonymousFunction()

	people := []chapter5.Person{
		{"Sheldon","Cooper"},
		{"Tom","Hardy"},
		{"Hugh","Jackman"},
		{"Tom","Hanks"},
		{"Ben","Big"},
	}
	chapter5.TestClosure(people)
}
