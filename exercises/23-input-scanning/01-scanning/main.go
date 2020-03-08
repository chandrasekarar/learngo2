package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fmt.Printf("Scanned Input %s\n", in.Text())
	}

	//file := os.Stdin

}