package main

import (
	"fmt"
	"os"
	"os/user"

	"Gengo/repl"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Gengo programming language!\n", usr.Username)
	fmt.Print("Feel free to type in commands\n")
	repl.StartVM(os.Stdin, os.Stdout)
}
