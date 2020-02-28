package main

import "fmt"

func main() {
	var (
		names    []string
		distance []int
		data     []uint8
		ratios   []float64
		alives   []bool
	)

	names = []string{}
	distance = []int{}
	data = []uint8{}
	ratios = []float64{}
	alives = []bool{}

	fmt.Printf("Length %d and Type %T of names\n",
		len(names), names)
	fmt.Printf("Are equal to nil %t \n\n",
		names == nil)

	fmt.Printf("Length %d and Type %T of distance\n",
		len(distance), distance)
	fmt.Printf("Are equal to nil %t \n\n",
		distance == nil)

	fmt.Printf("Length %d and Type %T of data\n",
		len(data), data)
	fmt.Printf("Are equal to nil %t \n\n",
		data == nil)

	fmt.Printf("Length %d and Type %T of ratios\n",
		len(ratios), ratios)
	fmt.Printf("Are equal to nil %t \n\n",
		ratios == nil)

	fmt.Printf("Length %d and Type %T of names\n",
		len(alives), alives)
	fmt.Printf("Are equal to nil %t \n\n",
		alives == nil)
}
