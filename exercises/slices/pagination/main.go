package main

import (
	"fmt"

	ps "github.com/inancgumus/prettyslice"
)

func main() {
	var movies = []string{
		"Bahubali", "Kabali", "Kurukshetra", "Majili",
		"Pilwan", "The Lion King", "Adi Lakshmi Purans",
		"Darbar", "Manikarnika", "Pokiri", "Maharshi",
		"Venky Mama", "Tubelight",
	}

	const pageSize = 4
	for from := 0; from < len(movies); from += pageSize {
		to := from + pageSize

		if to > len(movies) {
			to = len(movies)
		}

		currentPage := movies[from:to]
		head := fmt.Sprintf("Page #%d", (from/pageSize)+1)

		ps.Show(head, currentPage)
	}
}
