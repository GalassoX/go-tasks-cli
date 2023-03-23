package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/GalassoX/go-cli-tasks/models"
	"github.com/GalassoX/go-cli-tasks/utils"
)

func AddTask(input string, file *os.File, tasks *[]models.Task) bool {
	if len(input) < 1 {
		utils.Color(utils.CYAN, "Uso: crear [descripción de la tarea]")
		return false
	}

	tasksValue := *tasks

	var newID int
	if len(tasksValue) == 0 {
		newID = 1
	} else {
		newID = tasksValue[len(tasksValue)-1].ID + 1
	}

	task := models.Task{
		ID:        newID,
		Text:      input,
		Completed: false,
	}

	tasksValue = append(tasksValue, task)
	*tasks = tasksValue
	saveTask(file, *tasks)

	utils.Color(utils.GREEN, "\nTarea creada.\n\n")

	return true
}

func ListTasks(tasks []models.Task) bool {
	utils.Color(utils.YELLOW, "--|\t|-----------|\t|----------")
	fmt.Println("ID|\t|Descripción|\t|¿Completado?")
	utils.Color(utils.YELLOW, "--|\t|-----------|\t|----------")
	if len(tasks) > 0 {
		for _, task := range tasks {
			finished := "No"
			if task.Completed {
				finished = "Sí"
			}
			fmt.Printf("%d|\t|%s|\t|%s\n", task.ID, task.Text, finished)
		}
	} else {
		fmt.Println("No hay tareas para mostrar")
	}
	utils.Color(utils.YELLOW, "--|\t|-----------|\t|----------")
	return true
}

func DeleteTask(input string, file *os.File, tasks *[]models.Task) bool {
	if len(input) < 1 {
		utils.Color(utils.CYAN, "Uso: borrar [tarea id]")
		return false
	}

	id, err := strconv.ParseInt(input, 0, 64)
	if err != nil {
		utils.Color(utils.RED, "Error: Solo son admitidos numeros enteros")
		return false
	}

	tasksValue := *tasks

	for i, task := range tasksValue {
		if task.ID == int(id) {
			tasksValue = append(tasksValue[:i], tasksValue[i+1:]...)
			*tasks = tasksValue
			saveTask(file, *tasks)
			fmt.Print("\nTarea eliminada\n\n")
			return true
		}
	}
	utils.Color(utils.RED, "\nNo se encontró esa tarea\n\n")
	return false
}

func CompleteTask(input string, file *os.File, tasks *[]models.Task) bool {
	if len(input) < 1 {
		utils.Color(utils.CYAN, "Uso: completar [id tarea]")
		return false
	}

	id, err := strconv.ParseInt(input, 0, 64)
	if err != nil {
		utils.Color(utils.RED, "Error: Solo son admitidos numeros enteros")
		return false
	}

	tasksValue := *tasks

	for i, task := range tasksValue {
		if task.ID == int(id) {
			tasksValue[i].Completed = true
			*tasks = tasksValue

			saveTask(file, *tasks)
			fmt.Print("\nTarea lista!\n\n")
			return true
		}
	}
	utils.Color(utils.RED, "\nNo se encontró esa tarea\n\n")
	return false
}

func DeleteAllTasks(file *os.File, tasks *[]models.Task) bool {
	*tasks = []models.Task{}
	saveTask(file, *tasks)
	utils.Color(utils.GREEN, "\nLista de tareas eliminada\n\n")
	return true
}

func saveTask(file *os.File, tasks []models.Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		panic(err)
	}

	if err := file.Truncate(0); err != nil {
		panic(err)
	}

	if _, err := file.Write(bytes); err != nil {
		panic(err)
	}
}
