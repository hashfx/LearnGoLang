package main

import (
	"fmt"
)

func main(){
	fmt.Println("Interfaces in GoLang")
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))


	// 
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i<10; i++{
		fmt.Println(inc.Increment())
	}


}


// defining interface
type Writer interface {
	Write([]byte) (int, error)  // accept a slice of bytes and returns integer or error
}

// implementing interface
type ConsoleWriter struct {}  // implicit implemented

// creating a methods on ConsoleWriter with a signature of writer interface
func (cw ConsoleWriter) Write(data []byte) (int, error) {  // accepts slice of byte; returns integer or an error
	n, err := fmt.Println(string(data))
	return n, err
}

// 
type Incrementer interface{
	Increment() int
}

//
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
