package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	rx := regexp.MustCompile(`[^a-z]+`)

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf("Provide the search word\n")
		return
	}
	query := args[0]

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	 words := make(map[string]bool)

	for in.Scan() {
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")
		if len(word) > 2 {
			words[word] = true
		}
	}

	fmt.Printf("Length of words map are %d\n", len(words))

	if words[query] {
		fmt.Printf("Query contains word %q\n", query)
		return
	}
	fmt.Printf("Given word is not available\n")
}
