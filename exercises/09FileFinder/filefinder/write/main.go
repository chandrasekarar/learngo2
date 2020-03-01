package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	// for benchmark
	start := time.Now()

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	var names []byte
	for _, file := range files {
		name := file.Name()

		if file.Size() == 0 {
			names = append(names, name...)
			names = append(names, '\n')
		}
	}

	if len(names) == 0 {
		return
	}

	err = ioutil.WriteFile("out.txt", names, 0744)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", names)

	//print elapsed time
	fmt.Printf("Time taken to execute program: %dns\n", time.Since(start).Nanoseconds())
}
