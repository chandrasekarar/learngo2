package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(args)
	var total int
	for _, arg := range args {
		total += len(arg) + 1
	}
	// comment below code
	fmt.Printf("allocated cap: %d\n", total)
	names := make([]byte, 0, total)

	for _, arg := range args {
		names = append(names, arg...)
		names = append(names, '\n')
	}

	// comment below code
	fmt.Printf("actual cap: %d\n", cap(names))

	err := ioutil.WriteFile("sorted.txt", names, 0744)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", names)
}
