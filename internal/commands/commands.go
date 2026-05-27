package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"todo/internal/models"
	"todo/internal/storange"
)

func noteAdd() {
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

func noteDelete() {
	for _, tasks := range tasks {
		fmt.Printf("%d -%S\n", task.ID, task.Title)
	}
	fmt.Println("Qual o id da tarefa: ")
	var id int
	fmt.Scanln(&id)
	for index, task := range tasks {
		if task.Id == id {
			tasks.append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
}

func noteComplete() {
}

func noteView() {
}

func noteEdit() {
}
