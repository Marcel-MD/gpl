package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Marcel-MD/gpl/repl"
)

func main() {

	if len(os.Args) > 1 {
		path := os.Args[1]
		repl.ExecuteFile(path, os.Stdout)
		return
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands:\n")
	repl.Start(os.Stdin, os.Stdout)
}
