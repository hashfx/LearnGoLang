package main

import (
	"fmt"
)

func main() {

	num := 18
	fmt.Print("Guess a number: ")
	var guess int // guessed number stored in variable named guess
	fmt.Scan(&guess)
	// check if user enters number lesser than 1 or more than 100
	if guess < 1 || guess > 100{  // OR operator 
		fmt.Println("Number should be between 1 and 100!")
	}
	if guess >=1 && guess <= 100{  // AND operator
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

}
