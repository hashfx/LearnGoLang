/*

if Statement
	if condition {
		code to be executed
	}

if...else...if Statement
	if condition {
		code to be executed
	}
	else if {
		code to be executed
	}

if...else Statement
	if condition {
		code to be executed
	}
	else if {
		code to be executed
	}

Note:
	No single line blocks are allowed to execute in GoLang
	If block without '{' and '}' would return a Syntax Error
	if condition
		code


*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Control Flow in GoLang!")

	// a simple if statement, executes iff when condition is true
	if true {
		fmt.Println("Test is true")
	}

	// Comparison Operation using if ladder
	num := 18
	fmt.Print("Guess a number: ")
	var guess int // guessed number stored in variable named guess
	fmt.Scan(&guess)
	// check if user enters number lesser than 1 or more than 100
	if guess < 1{
		fmt.Println("Number should be greater than 1!")
	} else if guess > 100 {
		fmt.Println("Number should be smaller than 100!")

	} else {
		// above code is same as : if guess < 1 || guess > 100 { // code }  // OR operator 
	// if guess >=1 && guess <= 100{  // AND operator
		if guess < num { // if guessed number is less than number itself
			fmt.Println("Too Low")
		}
		if guess > num { // if guessed number is greater than number itself
			fmt.Println("Too High")
		}
		if guess == num { // if guessed number is equal to number itself
			fmt.Println("You guessed it right!")
		}
	}

	// returns true or false based on user-input
	fmt.Printf("number<=guess: %v\nnumber>=guess: %v\nnumber!=guess: %v\nnumber==guess: %v\n", num <= guess, num >= guess, num != guess, num == guess)

	// working with floats in GoLang Control Flow
	num2 := 0.123456789

	// | [number / (number^2)]-1 | < 0.001 : because floating point numbers are approximate estimation of decimal values
	if math.Abs(num2 / math.Pow(math.Sqrt(num2), 2)-1) < 0.001 {
		fmt.Println("They are same!")
	} else {
		fmt.Println("They are different!")
	}

}
