package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("use: todo [add|view|delete|completed|edit]")
		return
	}
	command := os.Args[1]

	switch command {

	case "add":
		commands.noteAdd()

	case "delete":
		commands.noteDelete()

	case "comp":
		commands.noteComplete()

	case "view":
		commands.noteView()

	case "edit":
		commands.noteEdit()

	}
}
