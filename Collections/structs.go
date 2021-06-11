package main

import (
	"fmt"
	"reflect"  // Reflection package for tags :: `required max: 100`
)

// struct
type Profession struct {
	id     int
	name   string
	skills []string // slice of type string
}

/* for Embedded purpose*/

type Animal struct{
	Name string
	Origin string
}
type Bird struct{
	// embedded struct Animal to struct Bird
	Animal  // struct Bird has inherited properties from struct Animal
	SpeedKPH float32
	CanFly bool
}

/* for Tags purpose */

type Validate struct{
	Name string `required max: "100"`  // Name field is required && maximum length of string can be '100'
	Grade int
}

func main() {
	fmt.Println("Structures in GoLang")

	Developer := Profession{
		id:   101,
		name: "myName",
		skills: []string{ // slice of type string
			"Programming",
			"Writing",
			"Critical Thinking",
		},
		/* Same output as above using [positional syntax] :: not recommended to use; can cause error if more elements in struct are added
		101,
		"myName",
		[]string{
			"Programming",
			"Writing",
			"Critical Thinking",
		},
		*/
	}

	// displaying single members of structure
	fmt.Printf("ID: %v\n", Developer.id)
	fmt.Printf("Name: %v\n", Developer.name)
	fmt.Printf("Skills: %v\n", Developer.skills[2])  // printing value at index[2] from slice in structure

	// structures with single element can be declared as follow :: Anonymous Structure
	Programmer := struct{name string}{name: "Bill Gates"}
	ProgrammerTwo := Programmer
	// unlike slices and maps, struct refers to independent datasets
	ProgrammerTwo.name = "Steve Jobs"  // changed 'name' to "Steve Jobs"
	fmt.Println(Programmer)  // {Bill Gates}
	fmt.Println(ProgrammerTwo)  // {Steve Jobs}
	
	// Embedded in Structures

	// struct Animal and struct Bird are still independent structures and have no relationship
	// Bird has Animal like characteristics but it not same as Animal :: they can't be used interchangeably
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	/* above code is same as
	b := Bird{
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPM: 48,
		CanFly: false,
	}
	*/
	fmt.Println(b)


	/* Tags in Struct */

	t := reflect.TypeOf(Validate{})  // get TypeOf object using reflection package
	field, _ := t.FieldByName("Name")  // get tag for "Name" field
	fmt.Println(field.Tag) // required max: "100"



}
