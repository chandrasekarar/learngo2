package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	max, _ := strconv.Atoi(os.Args[1])

	var uniques []int

loop:
	for len(uniques) <= max {
		r := rand.Intn(max + 1)
		fmt.Print(r, " ")

		for _, u := range uniques {
			if u == r {
				continue loop
			}
		}

		uniques = append(uniques, r)
	}

	fmt.Println("\n\nuniquies: ", uniques)

	sort.Ints(uniques)

	fmt.Println("\n\nuniquies\n: ", uniques)

	// var books [5]string
	// books[0] = "dracula"
	// books[1] = "1984"
	// books[2] = "island"
	// books[3] = "duper"

	// newBooks := [...]string{
	// 	"Go Lang",
	// 	"Go In Action",
	// }

	// games := []string{"arera", "zerara"}

	// fmt.Printf("Books		:%#v\n", books)

	// fmt.Printf("New Books		:%#v\n", newBooks)

	// fmt.Printf("Games		:%#v\n", games)
	// fmt.Printf("Games Length	:%#v\n", len(games))
	// fmt.Printf("Games		:%#v\n", games[0])
}
