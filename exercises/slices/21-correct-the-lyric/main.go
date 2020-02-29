package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Correct the lyric
//
//  You have a slice that contains the words of Beatles'
//  legendary song: Yesterday. However, the order of the
//  words are incorrect.
//
// CURRENT OUTPUT
//
//  [all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
// EXPECTED OUTPUT
//
//  [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
//
// STEPS
//
//  INITIAL SLICE:
//    [all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
//
//  1. Prepend "yesterday" to the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
//
//  2. Put the words to the correct positions in the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
//
//  3. Print the `lyric` slice.
//
//
// BONUS
//
//   + Think about when does the append allocate a new backing array.
//
//   + Check whether your conclusions are correct.
//
//
// HINTS
//
//   If you get stuck, check out the hints.md file.
//
// ---------------------------------------------------------

func main() {
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)

	// ADD YOUR CODE BELOW:
	// ...
	lyric = append([]string{"yesterday"}, lyric...)
	fmt.Printf("length of lyric: %d\n", len(lyric))

	const N, M = 8, 13

	lyric = append(lyric, lyric[N:M]...)
	lyric = append(lyric[:N], lyric[M:]...)

	fmt.Printf("\n%s\n\n", lyric)

	a, b := []int{1, 2, 3}, []int{4, 5, 6}

	c := append(a, b...)
	fmt.Printf("\n%d\n\n", c)

}
