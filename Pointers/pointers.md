# Pointers in GoLang

## Creating Pointers
  + ### Pointer types use an asterisk (**\***) as a prefix to type pointed to
    + ### **\*int** : a pointer to an integer
  + ### using **`addressof (&)`** operator to get address of variable

## Dereferencing Pointers
  + ### Dereference a pointer by preceding with an asterisk (*)
  + ### Complex types (such as structs) are automatically dereferenced

## Create pointers to objects
  + ### can use `addressof (&)` operator if value type already exists
    + ### `structure_object := structure_name{value}`
    + ### `pointer := &structure_object  // created pointer to structure`
  + ### use `addressof (&)` operator before initializer
    + ### `&structure_name{value}`
  + ### use the `new` keyword
    + ### can't initialize fields at the same time

## Types with internal pointers
  + ### All assignments operations in Go are copr operations
  + ### Slices and maps contain pointers, so copies point to same underlying data
