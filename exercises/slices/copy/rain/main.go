package main

import (
	"fmt"

	"github.com/inancgumus/prettyslice"
)

func main() {
	prettyslice.PrintBacking = true

	data := []float64{10, 25, 30, 50}
	prettyslice.Show("Probabilites", data)

	newData := []float64{99, 100}

	copy(data, newData) // it changes first 2 elements of data

	prettyslice.Show("Probabilites", data)

	fmt.Printf("Is it gonna rain? %.f%% chance\n",
		(data[0]+data[1]+data[2]+data[3])/float64(len(data)))

	// How to keep existing data and copy new
	data = []float64{10, 25, 30, 50}

	saved := make([]float64, len(data))
	copy(saved, data)
	//otherway use append to get backup
	saved2 := append([]float64(nil), data...)
	prettyslice.Show("Probabilites Saved", saved)
	prettyslice.Show("Probabilites Saved2", saved2)
	//lets update data
	// notice saved and saved2 not changed
	copy(data, newData)
	prettyslice.Show("Probabilites data", data)
	prettyslice.Show("Probabilites Saved", saved)
	prettyslice.Show("Probabilites Saved2", saved2)

	nilSlice := []string(nil)
	prettyslice.Show("nilSlice", nilSlice)

	// empty slice creates pointer reference
	// always use nil slice to initialize
	emptySlice := []string{}
	prettyslice.Show("emptySlice", emptySlice)
}
