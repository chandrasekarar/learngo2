package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool
type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	u := []user{
		{"chan", "1234", nil},
		{"god", "108", permissions{"admin": true}},
		{"devil", "100", permissions{"write": true}},
	}

	out, err := json.MarshalIndent(u, "", "\t")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
