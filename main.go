package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/GalassoX/go-cli-tasks/internal"
	"github.com/GalassoX/go-cli-tasks/models"
)

func printActions() {
	fmt.Println("Opciones disponibles: [ver|crear|completar|borrar|cerrar|borrartodo|ayuda]")
}

func main() {
	cmd := exec.Command("cmd", "/c", "cls && color ")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("cmd", "/c", "color ")
	cmd.Run()

	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	var tasks []models.Task

	info, err := file.Stat()
	if err != nil {
		panic(err.Error())
	}

	if info.Size() > 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		if err = json.Unmarshal(bytes, &tasks); err != nil {
			panic(err)
		}
	} else {
		tasks = []models.Task{}
	}

	fmt.Println("Bienvenido!")
	printActions()

	close, showActions := false, false
	for !close {

		if showActions {
			printActions()
		}

		var input string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input += scanner.Text()
		}

		in := strings.Split(input, " ")
		cmd := in[0]
		args := strings.Join(append(in[:0], in[1:]...), " ")

		switch cmd {
		case "listar", "lista", "ver":
			showActions = internal.ListTasks(tasks)
		case "crear":
			showActions = internal.AddTask(args, file, &tasks)
		case "borrar":
			showActions = internal.DeleteTask(args, file, &tasks)
		case "completar":
			showActions = internal.CompleteTask(args, file, &tasks)
		case "limpiar":
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			showActions = true
		case "cerrar":
			close = true
		case "borrartodo", "eliminartodo":
			showActions = internal.DeleteAllTasks(file, &tasks)
		case "ayuda":
			showActions = true
		default:
			showActions = true
		}
	}
}
