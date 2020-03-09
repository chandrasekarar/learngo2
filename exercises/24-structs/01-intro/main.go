package main

import "fmt"

func main() {
	type person struct {
		name, lastname string
		age            int
	}

	cha := person{
		name:     "Chandrasekara",
		lastname: "Reddy",
		age:      31,
	}

	var neeru person

	neeru.name = "Neeraja"
	neeru.lastname = cha.name + " " + cha.lastname
	neeru.age = 22

	fmt.Printf("%+v\n", cha)
	fmt.Printf("%#v\n", neeru)

}
