package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kylidboy/monkey_interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
