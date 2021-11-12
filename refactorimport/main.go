package main

import (
	"fmt"
)

func main() {
	c := &Config{
		RepoPath: "test amend",
		Rule:     map[string]string{},
	}
	fmt.Println(c)
}
