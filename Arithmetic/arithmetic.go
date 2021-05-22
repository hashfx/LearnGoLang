package arithmetic

import "fmt"

func Arithmetic(){

	fmt.Println("Arithmetic Operations in GoLang: ")

	fmt.Println("100 + 10 = ", 100 + 10)
	fmt.Println("100 - 10 = ", 100 - 10)
	fmt.Println("100 * 10 = ", 100 * 10)
	fmt.Println("100 / 10 = ", 100 / 10)
	fmt.Println("100 % 10 = ", 100 % 10)

	const pi float32= 3.14
	var num float32 = 5.1
	area := pi * num * num
	fmt.Println(area)
}

func Strings(){
	var myName string = "My Name"
	fmt.Println("Length of String is: ", len(myName))
}

func datatype(){
	const pi = 3.14159
	fmt.Printf("Data type is: %T", pi)
}

func logic(){
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)
}

func loops(){
	i := 1
	
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("----------")
	
	for j := 0; j < 5; j++{
		fmt.Println(j)
	}
	fmt.Println("----------")
	
	fmt.Println("----------")
	
}

func main(){
	loops()
}