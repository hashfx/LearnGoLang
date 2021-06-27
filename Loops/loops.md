# Loops in GoLang

## `for` statements

  + ### simple loops
    + ### `for initializer, condition, incrementor{}`
      + ### iterates from initial initializer value to until condition is true; if condition is true; increases the initializer variable
    + ### `for condition {}`
      + ### same as while loop : exists loop when condition gets false
    + ### `for {}`
      + ### infinite loop : can be terminated using break keyword

  + ### exiting early from loop
    + ### break
      + ### exists the loop
    + ### continue
      + ### breaks out of current iteration but not the loop, goes back to incrementor and continues execution of loop at next incrementer
    + ### labels
      + ### breaks out of the labelled loop defined in a nested loop
  
  + ### looping over collections
    + ### works for arrays, slices, maps, strings, channels
    + ### `for key, value := range collection {}`

