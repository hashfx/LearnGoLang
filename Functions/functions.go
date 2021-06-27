/*
Syntax:
	func function_name(parameter type of parameter) return datatype {
		// code
	}
	Syntactic Sugar
		func function_name(parameter type) (return_variable_name int) {
			code
			return  // just return keyword and no other expression
		}

Naming Convention:
	Function name could be in camelCase or PascalCase notation
	Initial letter of name, being uppercase or lowercase, determines the visibility of that function
		Uppercase first letter is going to be published from the package
		lowercase first letter to be kept internal to the package

Importants:
	> Function should start with func keyword, followed by function name, matched parantheses, and opening { (curly brace) on the same line
	> { curly brace should be on the same line where function is declared
	> Parameters are treated just like local variables we create locally
	> When calling a function, parameters have to be in same order as declared.
	> If parameters are of same type, say string, type can be declared at last for once
		> str(str1, str2 string) { code }
	> Pass by Value: data is not going to change when passed in it
	> When working with pointers in function as parameter
		> When function is declared, it's actually declared on execution stack of the function
		> When this function is exited, then execution stack is destroyed and that's memory is freed up
		> It's not safe because you are returning a pointer to a location in memory that just got freed
		> In GoLang, when runtime recognizes that returning value is generated on the local stack,
			it automatically promotes that variable to be on shared memory of computer, aka Heap memory




*/

package main // entry point of Go application

import (
	"fmt"
)

func main() { // function main takes no parameter and returns no value
	fmt.Println("Functions in GoLang")

	greeting := "Hello"
	name := "Google"

	// calling greet() function with argument as Google
	fmt.Println("\nFunction 1 : greet()")
	greet(name) // Hello, Google!

	// function with more than 1 parameters
	fmt.Println("\nFunction 2 : greet_2()")
	greet_2(greeting, name) // Hello Google

	// function with parameter as pointers
	fmt.Println("\nFunction 3 : greet()_3")
	greet_3(greeting, name) // Hello Google \n Microsoft
	fmt.Println(name)       // google

	fmt.Println("\nFunction 4 : greet_4()")
	greet_4(&greeting, &name) // Hello, Google \n Microsoft
	fmt.Println(name)         // Microsoft

	// function with variable arguments
	fmt.Println("\nFunction 5 : Function with Variable Parameters")
	res := sum(10, 20, 30, 40, 50)
	fmt.Printf("The sum is: %v\n", res)
	// fmt.Printf("The sum is: %v", *res)  // 150 :: dereferencing pointer res

	// function with multiple return values
	fmt.Println("\nFunction with multiple return values")
	d, err := divide(5.0, 2.0) // +Inf :: err : initialized error parameter
	// d := divide(5.0, 0.0)  // (before error-return handling) +Inf :: Cannot divide by zero
	// check for error, if any
	if err != nil { // if error is occurred; division by zero
		fmt.Println(err) // display error
		return
	}
	fmt.Println(d)

	fmt.Println("\nWorking with Functions as Types")
	// Anonymous Function
	func() { // parameters are optional
		fmt.Println("Hello, Go! This is an Anonymous Function")
	}() // () invokes the function :: aka Immediately Invoked Function

	fmt.Println("\nUsing Anonymous function in a loop")
	// using Anonymous function in a loop
	for i := 0; i < 5; i++ {
		func(i int) { // using outer variable i in inner function
			fmt.Println(i)
		}(i) // invoking and calling function
	}

	// passing function as variable
	fmt.Println("\nPassing Anonymous Function as variable")
	// var f func() := func() { code }  // type signature
	f := func() { // declared an anonymous function and assigned it to variable f
		fmt.Println("Hello!")
	}
	f()

	// working with methods
	fmt.Println("\nworking with methods")
	g := greeter{ // g: struct object
		greeting: "Hello",
		name:     "Google",
	}
	g.greet()
	fmt.Println("The new name is: ", g.name) // The new name is: Google

}

// function to greet user with their name
func greet(name string) { // function parameter : name
	fmt.Println("Hello, " + name + "!")
}

// function with 2 parameters and same data type
func greet_2(greeting, name string) {
	fmt.Println(greeting, name)
}

// passing pointer as parameter to function
func greet_3(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Microsoft" // Microsoft
	fmt.Println(name)
}

// passing pointer as parameter to function
func greet_4(greeting, name *string) {
	fmt.Println(*greeting, *name) // dereferencing pointer
	*name = "Microsoft"
	fmt.Println(*name) // Microsoft
}

// Variatic Parameters
// func sum(values ...int) *int {  // returning pointer to integer
func sum(values ...int) int { // a function can have only one variatic parameter which should be declared at last if other parameters are also used
	fmt.Printf("Values: %v\n", values) // prints array with values
	result := 0                        // sum initialized to 0
	for _, v := range values {
		result += v // result = result + v
	}
	return result // returns final result obtained after loop is exited
	// return &result // returns address of the result
}

// returning multiple return values
func divide(a, b float64) (float64, error) { // 2 return values : (float64, error)
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by Zero")
	}
	return a / b, nil // return nil if no error is occurred
}

// Methods in GoLang
type greeter struct {
	greeting string
	name     string
}

// a method
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "New_Name" // name changed to empty string
}  // func (g *greeter) greet() {} returns New_Name (updated g.name) due to pointer *greeter
