package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file with their ordinals
//
//  Use the previous exercise.
//
//  This time, print the sorted items with ordinals
//  (see the expected output)
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
//     1. apple
//     2. banana
//     3. orange
//
//
// HINTS
//
//   ONLY READ THIS IF YOU GET STUCK
//
//   + You can use strconv.AppendInt function to append an int
//     to a byte slice. strconv contains a lot of functions for appending
//     other basic types to []byte slices as well.
//
//   + You can append individual characters to a byte slice using
//     rune literals (because: rune literal are typeless numerics):
//
//     var slice []byte
//     slice = append(slice, 'h', 'i', ' ', '!')
//     fmt.Printf("%s\n", slice)
//
//     Above code prints: hi !
// ---------------------------------------------------------

func main() {
	items := os.Args[1:]
	if len(items) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(items)
	var total int
	for _, s := range items {
		total += len(s) + 4
	}
	// comment below code
	fmt.Printf("allocated cap: %d\n", total)

	data := make([]byte, 0, total)
	for i, s := range items {
		data = strconv.AppendInt(data, int64(i+1), 10)
		data = append(data, '.', ' ')
		data = append(data, s...)
		data = append(data, '\n')
	}
	// comment below code
	fmt.Printf("actual cap: %d\n", cap(data))

	err := ioutil.WriteFile("sorted.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", data)
}
