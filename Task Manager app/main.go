package main

import (
	"bufio"
	"fmt"
	"local/commands"
	"os"
	"strings"
)

func main() {

	//Grettings
	fmt.Println("Welcome to Task Tracker!")
	fmt.Println("Please use the commands to manage your tasks. Once you're done, type 'exit'")

	userInput := bufio.NewScanner(os.Stdin)

	var tasklist []commands.Task

	if userInput.Scan() {

		input := userInput.Text()

		if strings.ToLower(input) != "exit" {
			var err error
			tasklist, err = commands.FiletoJson()

			if err != nil {
				fmt.Println("Error:", err)
				return
			}

		}

		for strings.ToLower(input) != "exit" {
			var err error
			tasklist, err = commands.ExecuteCommand(input, tasklist)

			if err != nil {
				fmt.Println("Error:", err)
			}

			userInput.Scan()
			input = userInput.Text()

			if strings.ToLower(input) == "exit" {
				err := commands.JsontoFile(tasklist)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
			}

		}

	} else if userInput.Err().Error() != "" {
		fmt.Println("Error while retrieving the command. Please try again.")
	}
}
