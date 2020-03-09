package main

import "fmt"

func main() {
	type text struct {
		title string
		words int
	}

	type book struct {
		text
		isbn string
	}

	moby := book{
		text: text{
			title: "moby dick",
			words: 206052,
		},
		isbn: "ANB-10-20-30",
	}

	moby.text.words = 1000
	moby.words++

	fmt.Printf("\n%s has %d words. (isbn: %s)\n\n",
		moby.title, moby.words, moby.isbn)
}
