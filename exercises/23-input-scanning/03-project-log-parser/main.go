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
	var (
		log map[string]int
		domains []string
		total int
		lines int
	)
	in := bufio.NewScanner(os.Stdin)
	log = make(map[string]int)

	for in.Scan() {
		lines ++
		fields := strings.Fields(in.Text())

		if len(fields) < 2 {
			fmt.Printf("wrong input %v (line# %d)\n", fields, lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Printf("wrong input %q (line# %d)\n", fields[1], lines)
			return
		}

		if _, ok := log[domain]; !ok {
			domains = append(domains, domain)
		}
		total += visits
		log[domain] += visits
	}

	// print output
	fmt.Printf("%-30s %s\n", "DOMAIN", "VISITS")
	fmt.Printf("%s\n", strings.Repeat("_", 45))

	sort.Strings(domains)
	for _, domain := range domains {
		fmt.Printf("%-30s %3d\n", domain, log[domain])
	}
	fmt.Printf("%s\n", strings.Repeat("_", 45))
	fmt.Printf("%-30s %3d\n", "TOTAL", total)

	// handle error from scanner
	if err := in.Err(); err != nil {
		fmt.Printf("> Err: %q\n", err)
		return
	}
}
