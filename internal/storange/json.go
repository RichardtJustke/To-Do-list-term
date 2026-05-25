package storange

import (
	"encoding/json"
	"os"

	"todo/internal/models"
)

func SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("data/tasks.json", data, 0o644)

	return err
}

func LoadTasks() ([]models.Task, error) {
	data, err := os.ReadFile("data/tasks.json")
	if err != nil {
		return nil, err
	}

	var tasks []models.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
