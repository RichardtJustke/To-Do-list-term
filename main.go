package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"todo/internal/storange"
)

func main() {
	tasks, err := storange.LoadTasks()
	if err != nil {
		fmt.Println("Erro ao carregar as tarefas", err)
		return
	}
	//title tarefa
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o nome d tarefa: ")
	title, _ :=reader.ReadString('\n')
	title, strings.TrimSpace(title)
	//description tarefa
	//completed
