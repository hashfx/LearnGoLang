package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers in GoLang")

	a := 42
	b := a // copies data from a and assigns to b :: not pointing to same memory
	fmt.Println(a, b)
	a = 43
	fmt.Println(a, b) // value of a changed but value of b stays as previous

	var c int = 18
	var d *int = &c                      // pointer initialized as d :: d points to memory address of a
	fmt.Printf("c: %v, &c: %v\n", c, d)  // printing c and memory address of c
	fmt.Printf("&c: %v, d: %v\n", &c, d) // checking if &c is same as d :: both are exactly same

	fmt.Println("\nDereferencing Pointers")

	fmt.Printf("c: %v, *d: %v\n", c, *d) // finds memory location where pointer is pointing to and pull the value back out
	fmt.Println("Changing value of c changes that of d also")
	c = 25
	fmt.Printf("c: %v, *d: %v\n", c, *d) // displays same value because c and d points at same underlying data
	fmt.Println("Changing value of *d changes that of c also")
	*d = 17
	fmt.Printf("c: %v, *d: %v\n", c, *d) // displays same value because c and d points at same underlying data

	fmt.Println("\nVariable handling when they're assigned to one another")
	m := [5]int{10, 20, 30, 40, 50}    // array initialized
	n := m                             // n pointing to array m
	fmt.Printf("m: %v; n: %v\n", m, n) //
	m[1] = 25                          // changing index[1] of m to 25
	fmt.Printf("m is updated but n is not")
	fmt.Printf("m: %v; n: %v\n", m, n) // printing updated array

	fmt.Println("In case of slicing, j and k both are updated")
	j := []int{100, 200, 300, 400, 500} // slice initialized :: slice contains pointer to first element  // todo
	k := j
	fmt.Printf("m: %v; n: %v\n", j, k) // printing updated array
	j[1] = 250
	fmt.Printf("m: %v; n: %v\n", j, k) // printing updated array

	fmt.Println("\nPointers Arithmetic")
	e := [5]int{1, 2, 3, 4, 5}
	f := &e[0] // pointer to first element of array
	g := &e[1] // pointer to second element of array
	// g := &e[1] - 8   // Go doesn't allow
	fmt.Printf("%v %p %p\n", e, f, g) // displays value of array e and address of f and g
	fmt.Println("No Pointer Arithmetic Operation in GoLang")

	fmt.Println("Pointer Type")
	var ms *myStruct                                     // pointer to structure
	fmt.Printf("Uninitialized Pointer Struct: %v\n", ms) // uninitialized pointer is initialized a nil value by Go Compiler
	ms = &myStruct{foo: 42}                              // &{42}
	fmt.Printf("&myStruct: %v\n", ms)

	// using new keyword
	ms = new(myStruct)
	(*ms).foo = 42 // dereferencing ms :: dereferencing operator has lower precedence than the dot operator
	// (*ms).foo = 42
	ms.foo = 42                                      // pointer ms is pointer to struct that has a field foo
	fmt.Printf("Struct using new keyword: %v\n", ms) // &42
	fmt.Printf("(*ms): %v\n", (*ms).foo)             // dereferencing ms  :: 42(*ms).foo = 42
	fmt.Printf("ms: %v\n", ms.foo)                   // dereferencing ms  :: 42

}

type myStruct struct {
	foo int
}
