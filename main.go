package main

import (
	"Interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", u.Username)
	fmt.Printf("Feel free to type in commands!\n")

	repl.Start(os.Stdin, os.Stdout)
}
