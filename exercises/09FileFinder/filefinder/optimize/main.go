package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
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

	var total int
	for _, file := range files {
		name := file.Name()
		if file.Size() == 0 {
			total += len(name) + 1
		}
	}
	fmt.Printf("capacity required: %d\n", total)

	names := make([]byte, 0, total)
	for _, file := range files {
		name := file.Name()
		if file.Size() == 0 {
			names = append(names, name...)
			names = append(names, '\n')
		}
	}

	fmt.Printf("total capacity of names: %d\n", cap(names))

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
