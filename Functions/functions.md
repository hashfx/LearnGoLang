# `functions()` in GoLang

## Basic Syntax
  + ### `func function_name() { code }`
    + ### Uppercase first letter _means_ function is going to be published and allowed to be executed from outside of the package
    + ### lowercase first letter keeps function internal to the package
    + ### opening curly `{` brace need to be on same line as the `func` keyword and `}` be on its own line

## Parameters
  + ### Parameters allow to pass data into function to influence how that function executes
  + ### Comma delimited list of variables and types
    + ### `func function_name(param1 type, param2 type)`
  + ### Parameters of same type list type once
    + ### `func function_name(param1, param2 type)`
  + ### When pointers are passed in, the function can change the value of the caller
    + ### This is always true for data of slices and maps
  + ### Use variadic parameters to send list of same types in
    + ### Must be last parameter
    + ### Received as a slice
    + ### `func function_name(paramV ...type)`

## Return values
  + ### Single return values just list types
    + ### `func function_name() type`
  + ### Multiple return values surrounded by parentheses
    + ### `func function_name() (return variable, return type)`
    + ### The `(result type, error)` paradigm is a very common  idiom
  + ### Can use named return values
    + ### Initializes returned variable
    + ### Return using return keyword on its own
  + ### Can return addresses of local variables
    + ### Automatically promoted from local memory `(stack)` to shared memory `(heap)`

## Anonymous Functions
  + ### functions don't have names if they are:
    + ### Immediately invoked
        ```go
        func() {  // anonymous function without a name
            // code
        } ()  // function invoked immediately
        ```
    + ### Assigned to a variable or passed as an argument to a function
        ```go
        f := func() {  // assigned func() to variable f
            // code
        }
        f()  // function f invoked and called
        // function can only be invoked after it has been declared
        ```

## Functions as types
  + ### Can assign functions to variables or use as arguments and return values in functions
  + ### Type signature is like function signature. with no parameter names
    + ### `var f func(string, string, int) (int, error)`

## Methods
  + ### Function that executes in context of a type
  + ### format
    ```go
    func (receiver_variable method_type) function() {
        // code
    }
    ```
  + ### Receiver can be value or pointer
    + ### Value receiver gets copy of type
    + ### Pointer receiver gets pointer to type
