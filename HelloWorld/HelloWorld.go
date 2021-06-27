//
package main

import "fmt" // format package -- provides formatting for I/O -- pronunciated as 'fimt'
func main() { // main function
	fmt.Println("Hello, World!")
	/*
		Programme Output
		> Open CLI (Terminal, Command Prompt)
		> type: go run file_name.go  [press enter]
			> [In]  go run HelloWorld.go
			> [op]  Hello, World!
		To get information about any package or function:
		[In]	godoc package_name function_name
	*/

	// Take User Input in GoLang
	fmt.Print("Enter Your Name: ") // Input Statement
	var name string                // variable initialized to take as input
	fmt.Scan(&name)                // storing user-input in variable named name at its address

	fmt.Print("Hello, " + name + "!")  // display user input on screen


}
