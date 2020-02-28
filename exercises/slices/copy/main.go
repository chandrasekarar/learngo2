package main

import (
	"github.com/inancgumus/prettyslice"
)

func main() {
	evens := []int{2, 4}
	odds := []int{1, 3, 5}

	prettyslice.Show("evens", evens)
	prettyslice.Show("odds", odds)

	//lets copy odds to evens
	copiedNo := copy(evens, odds)
	prettyslice.Show("copiedNo", copiedNo)
	prettyslice.Show("evens", evens)

	//viceversa
	evens = []int{2, 4}

	//lets copy evens to odds
	copy(odds, evens)
	prettyslice.Show("odds", odds)
}
