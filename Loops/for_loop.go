/*
GoLang has only one type of loop: For Loop
Syntax
	for initializer; condition; increment/decrement{
		code
	}
The following code gives an error, because there is no commana operator in GoLang
	for i := 0, j := 5; i < 5; i++, j++ {  // i is scoped to for block
		fmt.Println(i)  // Output: error
	}
The following code, whereas, doesn't give error; because many variable can be instantiated at once in GoLang
	for i, j := 5, 0; i < 5; i, j = i+1, j+2 {  // i++, j++ throws an error because in GoLang, increment operator is not an expression
		fmt.Println(i, j)  // runs successfully
	}

All three statements inside for loop are required (mandatory) :: but a semi-colon is required in lieu for correct format
	i := 0  // i is scoped to the main function
	for ; i<5; {  // i is already initialized  :: same as for i<5 {code}
		code
		i++  // i is incremented
	}

do...while in GoLang
	for i<5 {  // same as for ; i<5; { code }
		code
		i++
	}

Infinite for...loop
	for...loop block is left empty and certain logic is given or break keyword is used when loop needs to be terminated
	for {
		code
		some condition where loop is to be terminated  // optional in case of Infinite Loop
	}

Using label in Loop
	Labels are used to
	Loop:  // label defined
		for p := 1; p <=3; p++{
			for q := 1; q <= 3; q++{
				fmt.Println(p * q)  // prints permutation of (i * j)
				if i*j <=3 {
					break Loop  // out of main for loop
				}
			}
		}

for...range Loop
 with all collection of GoLang
Syntax
	for key, value := range collection_name {
		code
	}

	


*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in GoLang")

	fmt.Println("Counting using For Loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("Variables can be instantiated at once in loops")
	for a, b := 0, 0; a < 5; a, b = a+1, b+2 { // i++, j++ throws an error because in GoLang, increment operator is not an expression
		fmt.Println(a, b) // runs successfully
	}

	fmt.Println("Bifurcate odd and even numbers using loops")
	num := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Array: %v\n", num)
	// var arr [10]int
	earr := []int{}
	oarr := []int{}

	for k := 0; k < len(num); k++ {
		if k%2 == 0 {
			earr = append(earr, k)
		}
		// else { oarr = append(oarr, k) }
		if k%2 == 0 { // odd numbers using continue keyword
			continue // skips even numbers
		}
		oarr = append(oarr, k)

	}

	fmt.Printf("Even: %v; Odd: %v\n", earr, oarr)

	fmt.Println("do...while loop in GoLang")

	n := 0
	for n < 5 {
		fmt.Println(n)
		n++
	}

	fmt.Println("Infinite for loop with break keyword")
	for {
		fmt.Println(n)
		n++
		if n == 10 { // a condition where loop may come to an end
			break // loop breaks when n is reached the given logic
		}
	}

	fmt.Println("Nested Loop")
	for p := 1; p <= 3; p++ {
		for q := 1; q <= 3; q++ {
			fmt.Println(p * q) // prints permutation of (i * j)
		}
	}

	fmt.Println("Labels in Loop")
Loop: // label defined for initial for...loop with variable
	for c := 1; c <= 3; c++ {
		for d := 1; d <= 3; d++ {
			fmt.Println(c * d) // prints permutation of (i * j)
			if c*d >= 3 {
				break Loop // out of main/initial/labelled for loop
			}
		}
	}

	fmt.Println("Working with Collections using for...range loop")
	s := []int{1, 2, 3, 4, 5} // slice
	fmt.Printf("Slice: %v\n", s)
	for key, value := range s {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	fmt.Println("Working with Strings using for...range loop")
	str := "Hello Go!"
	for _, v := range str {  // ignores key attribute
		// k: prints index of char in str; v: prints ASCII value of each character; string(v): prints characters of the string
		// fmt.Println(k, string(v), v)
		fmt.Println(string(v), v)  // when key attribute is ignored
	}



}
