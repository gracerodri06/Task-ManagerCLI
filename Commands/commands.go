package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
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

func ExecuteCommand(command string) error {

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

	tasklist, err := filetoJson()

	if err != nil {
		return err
	}

	switch action {
	case "add":
		var taskid string
		tasklist, taskid, err = addTask(actionArgs, tasklist)

		if err == nil {
			fmt.Printf("Task added successfully (ID: %s) \n", taskid)
		}

	case "update":
		tasklist, err = updateTask(actionArgs, tasklist)

		if err == nil {
			fmt.Println("Task updated successfully")
		}

	case "delete":
		tasklist, err = deleteTask(actionArgs, tasklist)

		if err == nil {
			fmt.Printf("Task ID %s deleted successfully \n", actionArgs)
		}

	case "mark-in-progress":
		tasklist, err = markInProgress(actionArgs, tasklist)

		if err == nil {
			fmt.Printf("Task ID %s marked as in-progress \n", actionArgs)
		}

	case "mark-done":
		tasklist, err = markDone(actionArgs, tasklist)

		if err == nil {
			fmt.Printf("Task ID %s marked as done \n", actionArgs)
		}

	case "list":
		err = list(actionArgs, tasklist)

	default:
		return errors.New("Invalid command")

	}

	if err != nil {
		return err
	}

	return JsontoFile(tasklist)
}

func addTask(args string, tasklist []Task) (tasks []Task, taskid string, err error) {

	if len(args) == 0 {
		return tasklist, "", errors.New("Invalid arguments for add command.")
	}

	lenTasks := len(tasklist)

	taskID := lenTasks + 1

	var newTask Task
	newTask.Id = strconv.Itoa(taskID)
	newTask.Description = args
	newTask.Status = "todo"
	newTask.CreatedAt = time.Now().Format(time.RFC3339)
	newTask.UpdatedAt = "0"
	return append(tasklist, newTask), newTask.Id, nil

}

func updateTask(args string, tasklist []Task) ([]Task, error) {

	updArgs := strings.SplitAfterN(args, " ", 2)

	if len(updArgs) != 2 {
		return tasklist, errors.New("Invalid arguments for update command.")
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
		return tasklist, errors.New("Invalid ID. Task not found.")
	}

	//Change the status
	record.Description = updArgs[1]
	record.UpdatedAt = time.Now().Format(time.RFC3339)

	tasklist = slices.Replace(tasklist, idx, idx+1, record)

	return tasklist, nil

}

func deleteTask(args string, tasklist []Task) ([]Task, error) {

	if len(args) == 0 {
		return tasklist, errors.New("Invalid arguments for delete command.")
	}

	initialLen := len(tasklist)

	tasklist = slices.DeleteFunc(tasklist, func(task Task) bool {
		return task.Id == args
	})

	finalLen := len(tasklist)

	if initialLen != finalLen {
		return tasklist, nil
	} else {
		return tasklist, errors.New("Was not possible to delete this task. Please check the ID.")
	}

}

func markInProgress(args string, tasklist []Task) ([]Task, error) {

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
		return tasklist, errors.New("Invalid ID. Task not found.")
	}

	//Change the status
	record.Status = "in-progress"
	record.UpdatedAt = time.Now().Format(time.RFC3339)

	tasklist = slices.Replace(tasklist, idx, idx+1, record)

	return tasklist, nil

}

func markDone(args string, tasklist []Task) ([]Task, error) {

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
		return tasklist, errors.New("Invalid ID. Task not found.")
	}

	//Change the status
	record.Status = "done"
	record.UpdatedAt = time.Now().Format(time.RFC3339)

	tasklist = slices.Replace(tasklist, idx, idx+1, record)

	return tasklist, nil

}

func list(args string, tasklist []Task) error {

	if args != "" && args != "in-progress" && args != "done" && args != "todo" {
		return errors.New("Invalid arguments for list command.")
	}

	if len(tasklist) == 0 {
		fmt.Println("No tasks defined.")
		return nil
	}

	for _, val := range tasklist {

		if args == "" || val.Status == args {
			fmt.Println(val)
		}
	}

	return nil

}

func filetoJson() ([]Task, error) {

	//Open the Json file
	file, err := os.Open("Task Tracker.json")

	if err != nil {
		file, err = os.Create("Task Tracker.json")

		if err != nil {
			return nil, fmt.Errorf("Error on file creation: %w", err)

		} else {
			//New file - return an empty slice
			return make([]Task, 0), nil
		}
	}

	defer func() {
		file.Close()
	}()

	dec := json.NewDecoder(file)

	_, err = dec.Token()
	if err != nil {
		return nil, fmt.Errorf("Error while reading the file: : %w", err)
	}

	tasks := make([]Task, 0)

	for dec.More() {
		var data Task
		if err := dec.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("Error while reading the file: %w", err)
		}

		tasks = append(tasks, data)
	}

	_, err = dec.Token()
	if err != nil {
		return nil, fmt.Errorf("Error while reading the file: %w", err)
	}

	return tasks, nil

}

func JsontoFile(tasks []Task) error {

	b, err := json.MarshalIndent(tasks, "", "")

	if err != nil {
		return fmt.Errorf("Error on file creation/opening: %w", err)
	}

	os.WriteFile("Task Tracker.json", b, 0644)

	return nil

}
