package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

type Task struct {
	Id          string
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

func ExecuteCommand(command string) {

	var action string
	var actionArgs string

	res := strings.SplitAfterN(command, " ", 2) // Action + Action argumments

	switch len(res) {

	case 1:
		action = strings.TrimSpace(res[0])
	case 2:
		action = strings.TrimSpace(res[0])
		actionArgs = res[1]
	}

	taskList := filetoJson()

	switch action {
	case "add":
		taskList = addTask(actionArgs, taskList)

	case "update":
		taskList = updateTask(actionArgs, taskList)

	case "delete":
		deleteTask(actionArgs, taskList)

	case "mark-in-progress":
		markInProgress(actionArgs, taskList)

	case "mark-done":
		markDone(actionArgs, taskList)

	case "list":
		list(actionArgs, taskList)

	default:
		fmt.Println("Invalid command")
		return

	}

	filetoJson()

}

func addTask(args string, taskList []Task) []Task {

	addArgs := strings.SplitAfterN(args, " ", 2)

	if len(addArgs) == 1 || len(addArgs) > 2 {
		fmt.Println("Error: Invalid arguments for add command.")
		return taskList
	}

	var newTask Task
	newTask.Id = strings.TrimSpace(addArgs[0])
	newTask.Description = addArgs[1]
	newTask.Status = "To-do"
	newTask.CreatedAt = time.Now().Format(time.RFC3339)
	newTask.UpdatedAt = "0"
	return append(taskList, newTask)

}

func updateTask(args string, tasklist []Task) []Task {

	updArgs := strings.SplitAfterN(args, " ", 2)

	if len(updArgs) != 2 {
		fmt.Println("Error: Invalid arguments for update command.")
	}

	idx := -1
	var record Task

	//Get the task from its ID
	for i, value := range tasklist {
		if value.Id == strings.TrimSpace(updArgs[0]) {
			idx = i
			record = value
			break
		}
	}

	if idx == -1 {
		fmt.Println("Error: Invalid ID. Task not found")
	}

	//Change the status
	record.Description = updArgs[1]
	record.UpdatedAt = time.Now().Format(time.RFC3339)

	tasklist = slices.Replace(tasklist, idx, idx+1, record)

	for _, val := range tasklist {
		fmt.Printf("%s: %s\n", val.Id, val.Description)
	}

	return tasklist

}

func deleteTask(args string, tasklist []Task) []Task {

	if len(args) == 0 {
		fmt.Println("Error: Invalid arguments for delete command.")
		return tasklist
	}

	//Pensar em alguma maneira de verificar que o delete rolou. Ou que o ID é inválido

	tasklist = slices.DeleteFunc(tasklist, func(task Task) bool {
		return task.Id == args
	})

	for _, val := range tasklist {
		fmt.Printf("%s: %s\n", val.Id, val.Status)
	}

	return tasklist
}

func markInProgress(args string, tasklist []Task) []Task {

	idx := -1
	var record Task

	//Get the task from its ID
	for i, value := range tasklist {
		if value.Id == args {
			idx = i
			record = value
			break
		}
	}

	if idx == -1 {
		fmt.Println("Error: Invalid ID. Task not found")
		return tasklist
	}

	//Change the status
	record.Status = "In Progress"
	record.UpdatedAt = time.Now().Format(time.RFC3339)

	tasklist = slices.Replace(tasklist, idx, idx+1, record)

	for _, val := range tasklist {
		fmt.Printf("%s: %s\n", val.Id, val.Status)
	}

	return tasklist

}

func markDone(args string, tasklist []Task) []Task {

	idx := -1
	var record Task

	//Get the task from its ID
	for i, value := range tasklist {
		if value.Id == args {
			idx = i
			record = value
			break
		}
	}

	if idx == -1 {
		fmt.Println("Error: Invalid ID. Task not found")
		return tasklist
	}

	//Change the status
	record.Status = "Done"
	record.UpdatedAt = time.Now().Format(time.RFC3339)

	tasklist = slices.Replace(tasklist, idx, idx+1, record)

	for _, val := range tasklist {
		fmt.Printf("%s: %s\n", val.Id, val.Status)
	}

	return tasklist

}

func list(args string, tasklist []Task) []Task {

	if args != " " && args != "in-progress" && args != "done" && args != "to-do" {
		fmt.Println("Error: Invalid arguments for list command.")
		return tasklist
	}

	if len(tasklist) == 0 {
		fmt.Println("No tasks defined.")
		return tasklist
	}

	for _, val := range tasklist {

		if args == "" || val.Status == args {
			fmt.Println(val)
		}
	}

	return tasklist

}

func filetoJson() []Task {

	//Open the Json file
	file, err := os.Open("Task Tracker.json")

	if err != nil {
		file, err = os.Create("Task Tracker.json")

		if err != nil {
			//error handling
		}
	}

	defer func() {
		file.Close()
	}()

	dec := json.NewDecoder(file)

	_, err = dec.Token()
	if err != nil {
		//fmt.Printf("%T: %v\n", t, t)
		log.Fatal(err)
	}

	tasks := make([]Task, 0)

	for dec.More() {
		var data Task
		if err := dec.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("%s", err)
			break
		}

		tasks = append(tasks, data)
	}

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	return tasks

}

func JsontoFile(tasks []Task) {

	b, _ := json.MarshalIndent(tasks, "", "")

	os.WriteFile("Task Tracker.json", b, 0644)

}
