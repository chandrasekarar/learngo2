package main

import (
	"fmt"

	ps "github.com/inancgumus/prettyslice"
)

func main() {
	var todo []string

	todo = append(todo, "sign")

	ps.Show("todo", todo)

	nums := []int{9, 7, 5}
	evens := append(nums, []int{2, 4, 6}...)

	fmt.Println(nums, evens)
}
