package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	spendings := fetch()

	for i, daily := range spendings {
		var total int
		for _, spending := range daily {
			total += spending
		}
		fmt.Printf("Day %d: %d\n", i+1, total)
	}
}

func fetch() [][]int {

	content := `10 20 45 70
30 30 80 10 78
14 04 56
70 100 32 132`
	lines := strings.Split(content, "\n")

	spendings := make([][]int, 0, len(lines))

	for _, line := range lines {

		fields := strings.Fields(line)
		spending := make([]int, 0, len(fields))

		for _, field := range fields {
			val, _ := strconv.Atoi(field)
			spending = append(spending, val)
		}
		spendings = append(spendings, spending)
	}
	return spendings
}
