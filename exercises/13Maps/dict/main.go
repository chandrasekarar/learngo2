package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Provide word")
		return
	}

	// eng to telugu dict
	dict := map[string]string{
		"Hello":     "Namaste",
		"Good":      "Bagundu",
		"Greetings": "Subakankshalu",
	}
	fmt.Printf("%s\n", dict[args[0]])

	//tel to eng dict

	telDict := make(map[string]string, len(dict))

	for k, v := range dict {
		telDict[v] = k
	}
	fmt.Printf("%s\n", telDict[args[0]])

	// delete
	delete(dict, "Good")
	fmt.Printf("%q\n", dict["Good"])

	// fully delete map
	// dict = nil // will not delete complete memory
	// below wii destroy fully
	for k := range dict {
		delete(dict, k)
	}

	fmt.Printf("Len of dict %d\n", len(dict))

}
