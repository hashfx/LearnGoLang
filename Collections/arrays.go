package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays in GoLang")

	// array_name := [size_of_array]datatype_of_array{values of array}

	marks := [5]int{10, 20, 30, 40, 50}
	cgpa := [...]int{10, 9, 8, 7, 9} // creates an array thats large enough to hold the data passed to array
	var students [3]string           // Empty array of size 3 and datatype string

	fmt.Printf("Marks: %v\n", marks)
	fmt.Printf("CGPA: %v\n", cgpa)
	fmt.Printf("Empty Array: %v\n", students)
	// assigning values to array students
	students[0] = "Alpha"
	students[1] = "Beta"
	students[2] = "Gamma"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student at index 1: %v\n", students[1])
	fmt.Printf("Number of Students: %v\n", len(students))

	// Array of Arrays
	fmt.Println("Identity Matrix")

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	// var identityMatrix [3][3]int = [3][3]int{ [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}  // same output as above code
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3}
	b := a // not copying array 'a' to 'b' but creating a different array 'b' with index[1] as 5
	// b := &a  // copies array 'a' to array 'b'; output: Array(a): [1 5 3] Array(b): [1 5 3]
	b[1] = 5 // changed 2nd element of array 'a' to 5
	fmt.Printf("Array(a): %v\n", a)
	fmt.Printf("Array(b): %v\n", b)

	/* Slice : covered in slices.go
	c := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// c := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}  // slicing would work for undeclared size arrays too
	d := c[:]                            // slice of all elements
	e := c[3:]                           // slice from 4th element to end
	f := c[:6]                           // slice first 6 elements
	g := c[3:6]                          // slice 4th, 5th, 6th element
	c[5] = 55                            // changing one element would affect all arrays
	fmt.Println(c)                       // [10 20 30 40 50 60 70 80 90 100]
	fmt.Printf("Length: %v\n", len(c))   // 10
	fmt.Printf("Capacity: %v\n", cap(c)) // 10
	fmt.Println(d)                       // [10 20 30 40 50 60 70 80 90 100]
	fmt.Println(e)                       // [40 50 60 70 80 90 100]
	fmt.Println(f)                       // [10 20 30 40 50 60]
	fmt.Println(g)                       // [40 50 60]

	fmt.Println("Slice Operations")

	k := []int{}
	// k := make([]int, 3, 100)  // make(data_type_of_array, size_of_array, capacity)
	fmt.Println(k)
	fmt.Printf("Length: %v\n", len(k))
	fmt.Printf("Capacity: %v\n", cap(k))
	k = append(k, 1)
	fmt.Println(k)
	fmt.Printf("Length: %v\n", len(k))
	fmt.Printf("Capacity: %v\n", cap(k))
	k = append(k, 2, 3, 4, 5)  // appending many elements in one go
	fmt.Println(k)
	fmt.Printf("Length: %v\n", len(k))
	fmt.Printf("Capacity: %v\n", cap(k))
	// A Slice can be appended to another slice using spread operator (...) operator
	k = append(k, []int{11, 12, 13, 14}...)  // appending many elements in one go
	fmt.Println(k)
	fmt.Printf("Length: %v\n", len(k))
	fmt.Printf("Capacity: %v\n", cap(k))

	// Stack Operations in Slice
	fmt.Println("Stack Operations in Slices")

	l := []int{90, 80, 70, 60, 50}
	// m := l[1:]  // removes 1st element from the slice  // [80 70 60 50]
	// n := l[:len(l)-1]  // removes last element from the slice  // [90 80 70 60]
	// o := append(l[:2], l[3:]...)  // remove an element from the middle by concatinating two slices together  // [90 80 60 50]
	fmt.Println(l)
	
	*/

}
