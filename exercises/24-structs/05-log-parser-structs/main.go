package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type results struct {
	domain string
	visits int
}

type log struct {
	sum   map[string]results
	keys  []string
	total int
	lines int
}

func main() {
	in := bufio.NewScanner(os.Stdin)

	log := log{
		sum: make(map[string]results),
	}

	for in.Scan() {
		log.lines++

		fields := strings.Fields(in.Text())
		if len(fields) < 2 {
			fmt.Printf("wrong input. %v (line #%d)\n", fields, log.lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])

		if err != nil || visits < 0 {
			fmt.Printf("wrong input. %v (line #%d)\n", fields[1], log.lines)
			return
		}

		if _, ok := log.sum[domain]; !ok {
			log.keys = append(log.keys, domain)
		}
		log.sum[domain] = results{
			domain: domain,
			visits: visits + log.sum[domain].visits,
		}
		log.total += visits
	}

	sort.Strings(log.keys)

	fmt.Printf("%-30s %10s\n", "DOMAINS", "VISITS")
	fmt.Printf("%s\n", strings.Repeat("-", 45))

	fmt.Printf("%v\n\n", log.keys)

	for _, v := range log.keys {
		fmt.Printf("%-30s %5d\n", v, log.sum[v].visits)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 45))
	fmt.Printf("%-30s %5d\n", "TOTAL", log.total)
}
