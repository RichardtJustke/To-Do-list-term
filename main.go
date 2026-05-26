package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"todo/internal/models"
	"todo/internal/storange"
)

func main() {
	tasks, err := storange.LoadTasks()
	if err != nil {
		fmt.Println("Erro ao carregar as tarefas", err)
		return
	}

	// escrever tarefa e salvar
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o nome d tarefa: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	newTask := models.Task{
		ID:        len(tasks) + 1,
		Title:     title,
		Completed: false,
	}
	tasks = append(tasks, newTask)
	err = storange.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Erro ao salvar:", err)
		return
	}
	fmt.Println("Tarefa adicionada com sucesso!")

	// escrever descrição da tarefa
}
