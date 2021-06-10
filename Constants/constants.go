package main

import (
	"fmt"
	// "math"  // module for Mathematical functions
)

const a int16 = 95 // constant declared at package level for shadowing purpose

// Enumerated Constants
const (
	// _ = iota  // if you don't want to assign 0 value, use _ (underscore symbol) 
	x = iota  // untyped constant with special symbol (or a counter) named iota
	y = iota
	z = iota
	m  // displays result without initialization because of an ongoing pattern
	n  // displays result without initialization because of an ongoing pattern
)
// iota is scoped to constant block
const x2 = iota
func main() {
	fmt.Println("Constants in GoLang!")
	// declaring constant
	const pi float32 = 3.14 // internal-constant in camelCase

	// const sinValue float32 = math.Sin(1)  // constants cannot be given values that are assign at runtime

	// pi = 3.145  // constant cannot be reassigned value :: Error: cannot assign to pi
	fmt.Printf("%v, %T\n", pi, pi)

	// constants can be of any primitive datatype
	const a int = 99          // integer constant  // commenting this line would display value of constant 'a' at package level
	const b string = "String" // string constant
	const c float32 = 3.14    // float constant
	const d bool = true       // boolean constant
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)

	// constants can be operated with variables
	const num1 int = 42
	var num2 int = 27

	// datatype decided by compiler itself
	const num3 = 37
	var num4 int16 = 49
	fmt.Printf("%v, %T\n", num1+num2, num1+num2)
	fmt.Printf("%v, %T\n", num3+num4, num3+num4)  // (num1+num2) same as (37+num2 : 37 being constant)

	/* Enumerated Constants declared at package level :: above main function */
	// value of constant iota is increased with each step
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", n)
	fmt.Printf("%v\n", x2)  // element of different constant block starting from 0
	

}
