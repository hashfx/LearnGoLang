# Arrays in GoLang

+ ### Collection of items with same type

+ ### fixed size

+ ### Declaration Style
  + ### `arr := [5]int{1, 2, 3, 4, 5}`
  + ### `arr := [...]int{1, 2, 3, 4, 5}`
  + ### `var arr [5]int`
+ ### Access via Zero-based index
  + ### `arr := [5]int{1, 2, 3, 4, 5} // arr[1] == 0`
  + ### **`len`** function returns size of array

# Slices in GoLang

  + ### Backed by array

  + ### Creation Styles
    + ### Slice existing array or slice

    + ### Literal style

    + ### via `make()` function

      + ### `a := make([]int, 10)`  // create slice with capacity and length = 10
      + ### `a := make([]int, 10, 100)`  // slice with length = 10 and capacity = 100

  + ### **`len`** function returns length of slice

  + ### **`cap`** function returns length of underlying array

  + ### **`append`** function to add elements to slice
    + ### may cause expensive copy operation if underlying array is too small
