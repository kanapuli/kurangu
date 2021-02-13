package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kanapuli/kurangu/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!!! This is Kurangu programming language\n", user.Username)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin)
}
