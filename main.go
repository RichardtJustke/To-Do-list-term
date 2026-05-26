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

	reader := bufio.NewReader(os.Stdin) // aqui lê o input do usuario não pode repetir!
	fmt.Print("Digite o nome da tarefa: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	// escrever descrição da tarefa
	fmt.Print("Digite a descrição da tarefa: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	newTask := models.Task{
		ID:          len(tasks) + 1,
		Title:       title,
		Description: description,
		Completed:   false,
	}
	tasks = append(tasks, newTask)
	err = storange.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Erro ao salvar:", err)
		return
	}
	fmt.Println("Tarefa adicionada com sucesso!")
}
