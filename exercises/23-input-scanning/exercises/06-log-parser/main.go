package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	var (
		domains map[string]int
		keys    []string
		total   int
		lines   int
	)
	domains = make(map[string]int)
	for in.Scan() {
		lines++
		fields := strings.Fields(in.Text())
		if len(fields) < 2 {
			fmt.Printf("Wrong input, missing field at (line %d)\n", lines)
			return
		}
		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Printf("Wrong input %v at (line %d)\n", fields[1], lines)
			return
		}
		if _, ok := domains[domain]; !ok {
			keys = append(keys, domain)
		}
		domains[domain] += visits
		total += visits
	}

	fmt.Printf("%-30s %10s\n", "Domains", "Visists")
	fmt.Printf("%s\n", strings.Repeat("-", 45))
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%-30s %5d\n", key, domains[key])
	}
	fmt.Printf("%s\n", strings.Repeat("-", 45))
	fmt.Printf("%-30s %5d\n", "Total", total)
}
