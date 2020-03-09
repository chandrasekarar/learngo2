package main

import (
	"bufio"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {
	// use your solution from the previous exercise
	fmt.Printf("To list games type (ls) and to quit type (q)\n")
	in := bufio.NewScanner(os.Stdin)
	games := []game{
		{item: item{1, "god of war", 50}, genre: "action adventure"},
		{item: item{2, "x-com 2", 30}, genre: "strategy"},
		{item: item{3, "minecraft", 20}, genre: "sandbox"},
	}

	_ = games

	for in.Scan() {
		switch in.Text() {
		case "ls":
			gameList(games)
		case "q":
			fmt.Printf("Visit again!!\n")
			return
		default:
			fmt.Printf("Invalid input\n")
			fmt.Printf("To list games type (ls) and to quit type (q)\n")
		}
	}
}

func gameList(games []game) {
	fmt.Println()
	fmt.Printf("Chandra`s game store has %d games.\n\n", len(games))
	for _, g := range games {
		fmt.Printf("#%d: %-15q %-20s $%d\n",
			g.id, g.name, "("+g.genre+")", g.price)
	}
	fmt.Println()
}
