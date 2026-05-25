package main

import (
	"fmt"

	"todo/internal/storange"
)

func main() {
	tasks, err := storange.LoadTasks()
	if err != nil {
		fmt.Println("Erro ao carregar as tarefas", err)
		return
	}
	fmt.Println(tasks)
}
