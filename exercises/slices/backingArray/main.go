package main

import (
	"fmt"
	"sort"

	ps "github.com/inancgumus/prettyslice"
)

func main() {
	grades := []float64{40, 10, 20, 50, 60, 70, 80}

	// front := append([]float64{}, grades[:3]...)
	front := grades[:3]

	ps.MaxPerLine = 7
	ps.PrintBacking = true
	ps.Show("Grades", grades)

	front[2] = 30

	sort.Float64s(front)

	ps.Show("Grades", grades)
	ps.Show("Front", front)

	// nums := []int{9, 7, 5, 3, 1}
	// nums = nums[:1]

	// ps.Show("nums", nums)

	// nums2 := nums[:]
	// ps.Show("nums2", nums2)

	arr := [...]int{9, 7, 5, 3, 1}
	nums := arr[2:] // [5, 3, 1]
	fmt.Println(nums)

	nums2 := nums[1:] // [3, 1]
	fmt.Println(nums2)

	arr[2]++          // arr[2] = arr[2] + 1
	fmt.Println(nums) // [6, 3, 1]

	fmt.Println("nums[1]", nums[1]) // 3
	fmt.Println("arr[4]", arr[4])   // 1

	nums[1] = nums[1] - (arr[4] - 4) // 3 - (-3) // 6
	fmt.Println("nums[1]", nums[1])

	fmt.Println(nums) // [6, 6, 1]
	nums2[1] += 5

	fmt.Println(nums) //[6, 6, 1]
}
