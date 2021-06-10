# Constants in GoLang

### Constants in GoLang are declared using `const` keyword

### **Naming Constants**: In GoLang, constants need not to be in uppercase (~~`MY_CONST`~~). Remember, if first letter of constant is in uppercase (My_const`), means the constant is going to be exported. So Constants will be named same as variables are declared 🙂

## Characteristics of Constants

+ ### Constants are Immutable and can be shadowed

+ ### Replaced by Compiler at Compile time

+ ### Constant has to be assignable at compile time :: functions are not acceptable as constants

  + ### Value must be calculable at Compile time

## Naming of Constants

  + ### Named like variables

  + ### PascalCase for **exported constants**

  + ### camelCase for **internal constants**

## Typed Constants

+ ### Typed constants work like immutable variables

+ ### can inter-operate only with same type

## Untyped Constants

+ ### Untyped constants work like literals

+ ### can inter-operate with similar types

## Enumerated Constants

+ ### Special symbol **`iota`** allows related constants to be created easily

+ ### `iota` starts at 0 in each `const` block and increments by one

+ ### watch out for constants values that match zero values for variables; they can cause bug in the program

## Enumerated expressions

+ ### Operations that can be determined at compile time are allowed

  + ### Arithmetic

  + ### BitWise operations

  + ### BitShifting
