/*
tests in switch case must not be overlapped
	switch condition {
	case 1:
		statement
	case 1:
		statement  // error: duplicate case switch
	case 2:
		statement
	}
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Control Flow in GoLang!")

	switch 2 { // if value of case is 2, return their respective statement
	case 1: // case with value 1
		fmt.Println("One")
		break  // stops control of switch case
		fmt.Println("This won't be displayed due to break statement")
	case 2: // case with value 2
		fmt.Println("Two")
	default: // default value: if both above cases are false, default would be displayed!
		fmt.Println("Nor one or two!")
	}

	// switch statement with multiple tests in single case
	switch 5 {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd")
	case 2, 4, 6, 8, 10:
		fmt.Println("Even")
	default:
		fmt.Println("NaN")
	}

	// Logical test cases in switch statement
	age := 16
	switch {
	case age < 16:
		fmt.Println("Not eligible for Driving License")
		// fallthrough  // executes statement of next case wheather next case is true or false
	case age >= 16 && age < 18:
		fmt.Println("Eligible for Learner's License")
	case age >= 18:
		fmt.Println("Eligible for Driving License")
	default:
		fmt.Println("Not eligible to run this Program!")
	}

	/* Type Switch */

	var i interface{} = []int{} // 1, 1.0, "1", []int{}
	switch i.(type) {
	case int:
		fmt.Println("i in an Integer")
	case float64:
		fmt.Println("i in a Float64")
	case string:
		fmt.Println("i in a String")
	case []int: // same array size and datatype must be used for this case to be true
		fmt.Println("i is an Array")
	default:
		fmt.Println("i is of another type")

	}

	

}
