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

	names = []string{"Neeraja", "Chandra", "Reddy"}
	distance = []int{120, 13, 40}
	data = []uint8{15, 72, 224}
	ratios = []float64{12.776, 98.6785, .542663}
	alives = []bool{true, true, false, true}

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

	var ok string
	if len(data) != len(distance) {
		ok = "not "
	}

	fmt.Printf("data and distance slice lengths are %sequal\n\n", ok)

	ok = ""
	if len(alives) != len(ratios) {
		ok = "not "
	}

	fmt.Printf("ratios and alives slice lengths are %sequal\n\n", ok)
}
