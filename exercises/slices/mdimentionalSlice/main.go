package main

import (
	"fmt"

	"github.com/inancgumus/prettyslice"
)

func main() {

	//test check type
	test := [][][3]int{{{10, 5, 9}}}
	fmt.Printf("Type %T \n", test)
	// Way 1
	spendings := [][]int{
		{200, 100},
		{500},
		{50, 25, 75},
	}

	// Way 2
	spendings = make([][]int, 0, 3)
	spendings = append(spendings, []int{200, 100})
	spendings = append(spendings, []int{500}, []int{50, 25, 75})

	// fmt.Printf("\nspendings%T%v\n\n", spendings, spendings)
	prettyslice.PrintBacking = true
	prettyslice.Show("spendings", spendings)
	for _, main := range spendings {
		// prettyslice.Show("day wise", main)
		sum := 0
		for _, inner := range main {
			// prettyslice.Show("each val", inner)
			sum += inner

		}
		prettyslice.Show("Daily Sum", sum)
	}
}
