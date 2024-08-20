package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/sudo-wave/cerc/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Cerc programming language!\n", user.Username)
	fmt.Printf("Please type in a command or multiple commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
