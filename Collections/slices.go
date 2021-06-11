package main
import(
	"fmt"
)

func main(){
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

	/* Stack Operations in Slice */
	fmt.Println("Stack Operations in Slices")

	l := []int{90, 80, 70, 60, 50}
	// m := l[1:]  // removes 1st element from the slice  // [80 70 60 50]
	// n := l[:len(l)-1]  // removes last element from the slice  // [90 80 70 60]
	// o := append(l[:2], l[3:]...)  // remove an element from the middle by concatinating two slices together  // [90 80 60 50]
	fmt.Println(l)
}	