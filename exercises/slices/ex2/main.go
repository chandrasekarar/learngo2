package main

import (
	"github.com/inancgumus/prettyslice"
)

func main() {
	prettyslice.PrintBacking = true

	nums := []int{1, 3, 2, 4}
	odds := nums[:2:2]
	odds = append(odds, 5, 7)

	evens := append(nums[2:4], 6, 8)
	prettyslice.Show("nums", nums)
	prettyslice.Show("odds", odds)
	prettyslice.Show("evens", evens)
}
