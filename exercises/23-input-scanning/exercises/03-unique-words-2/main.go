package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words 2
//
//  Use your solution from the previous "Unique Words"
//  exercise.
//
//  Before adding the words to your map, remove the
//  punctuation characters and numbers from them.
//
//
// BE CAREFUL
//
//  Now the shakespeare.txt contains upper and lower
//  case letters too.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < shakespeare.txt
//
//   There are 100 words, 69 of them are unique.
//
// ---------------------------------------------------------

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	var total int
	words := make(map[string]bool)
	rx := regexp.MustCompile(`[^A-Za-z]+`)

	for in.Scan() {
		total++
		word := rx.ReplaceAllString(in.Text(), "")
		word = strings.ToLower(word)
		words[word] = true
	}

	fmt.Printf("There are %d words and out off %d are unique\n", total, len(words))
}
