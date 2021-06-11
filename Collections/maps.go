package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")
	// map_name := map[datatype_of_key]datatype_of_values{key1:value1, key2:value2}
	/* creating map using make() function
	map_name := make(map[string]int, length_of_map(optional))
	map_name = map[string]int{map_values}
	   */
	Numbers := map[string]int{
		"Alpha":   01,
		"Beta":    10,
		"Gamma":   11,
		"Delta":   100,
		"Epsilon": 101,
	}
	fmt.Println(Numbers)
	
	/* Manipulating map values */
	
	// pulling one value from map by :=> map_name[value_name]
	fmt.Println(Numbers["Delta"])
	
	// adding elements in map
	Numbers["Zeta"] = 110  // map_name[value_to_be_added] = key_to_be_added
	fmt.Println(Numbers["Zeta"], Numbers)  // displaying added_value and whole_map
	
	// deleting elements in map
	delete(Numbers, "Zeta")  // delete(map_name, value_to_be_deleted)
	fmt.Println(Numbers)
	fmt.Println(Numbers["Zeta"])
	
	// 'comma-ok syntax' : returns true/false if element is found/not found in map
	pop, ok := Numbers["Epsion"]  // misspelt word Epsilon returns 0; return false if element is not in the map
	pop2, ok2 := Numbers["Epsilon"]  // returns key; return true if element is in the map
	fmt.Println(pop, ok)
	fmt.Println(pop2, ok2)

	// numbers of element in a map
	fmt.Println(len(Numbers))




}
