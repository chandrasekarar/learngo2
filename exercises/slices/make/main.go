package main

import (
	"fmt"
	"strings"

	"github.com/inancgumus/prettyslice"
)

func main() {
	prettyslice.PrintBacking = true
	prettyslice.MaxPerLine = 10

	// nums := make([]int, 3)
	// prettyslice.Show("make initialize", nums)

	// nums = make([]int, 3, 5)
	// prettyslice.Show("make cap increased", nums)

	tasks := []string{"jump", "run", "read"}

	var upTasks []string
	prettyslice.Show("upTasks", upTasks)
	for _, v := range tasks {
		//its inefficient, bcoz it allocates
		// 3 diffrent backing arrays
		upTasks = append(upTasks, strings.ToUpper(v))
		prettyslice.Show("upTasks", upTasks)
	}

	fmt.Println()

	//lets improve with make
	// it will assign an emptyy slice with
	// capacity 3
	upTasks = make([]string, 0, 3)
	prettyslice.Show("upTasks", upTasks)
	for _, v := range tasks {
		upTasks = append(upTasks, strings.ToUpper(v))
		prettyslice.Show("upTasks", upTasks)
	}

}
