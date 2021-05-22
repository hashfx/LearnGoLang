# Variables in GoLang

## **New** Variables can be declared in 3 ways in GoLang

### When variable is declared but not to be initialized {like in loops}

+ `var variable_name variable_datatype = value`
  + `var foo int = 40`

### To get control over datatype declaration along with variable

+ `variable_name = value`
  + `foo = 20`

### When full control is given to complier

+ `variable_name := value`
  + `foo = 40`

### Variables in GoLang are always have to be used. If not, it would return an error

## Declare many variables in one Go

```Go
var(
  developerID int = 101
  name string = "Google"
  pi float32 = 3.14
  grade string = "A+"
)
```

## Re-declaration of variable

+ Variable once declared cannot be re-declared but can be reassigned its value

```go
// Re-declaration of Variables
func main(){
  var number = 99
  // number := 100  // error: no new variables on left side of :=
  number = 100  // number is re-assigned a value
}
```

```go
var n int = 10
func main(){
  fmt.Println(n)  // Output: 10
  var n int = 25
  fmt.Println(n)  // Output: 25
}
```

## Naming of Variables

+ lower-case variables are scoped to the package (any file in the package can use that variable)
+ When a variable is declared inside a block (function/method), it is scoped to the block only (block: `func main()`)

## Naming convention in GoLang

+ Variable life is determined by its length
+ Acronyms should be kept upper-case for better readability (**myURL** and not **myUrl**)

## Typecasting of variables

### To Convert `int` to `string` without converting it to ASCII

```go
import (
  "fmt"
  "strconv" // library for typecast conversion
)

func main(){
  // Syntax: variable = datatype(to_be_casted_variable)
  var i int = 82
  // typecast int to float32
  var j float32
  j = float32(i)
  fmt.Println(%v, %T, j, j) // 82, float32

  // int to string

  var k string
  k = strconv.Itoa(k)
  fmt.Println("%v, %T", k, k)  // Output: 82, string 
}
```

```go
package main

import "fmt"

func main(){
    fmt.Println("Variables in GoLang")
}
```

# Summary

+ ## Variable Declaration
  + ### var foo int
  + ### var foo int = 10
  + ### foo := 10

+ ## Variables can't be re-declared, but can be shadowed  

+ ## All variables must be used

+ ## Visibility of Variables

  + ### Lower-case first letter for package scope
  + ### Upper-case first letter to export
  + ### There is no private scope, but can be scoped to a specific block
+ ## Naming Convention

  + ### Pascal or CamelCase
    + ### Capitalize acronyms (HTTP, URL)
  + ### As short as reasonable
    + ### Longer names for longer life

+ ## Type Conversion
  + ### DestinationType(variable)
  + ### use `strconv` package for strings
  + ### Go doesn't allow implicit type conversion
