package main

import "fmt"

type song struct {
	title  string
	artist string
}

type playlist struct {
	genre string
	songs []song
}

func main() {
	songs := []song{
		{title: "Hello", artist: "Mahesh"},
		{title: "Bombat", artist: "Nithin"},
	}

	rock := playlist{genre: "rocky", songs: songs}

	_ = rock
	fmt.Printf("%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}
}
