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

	if userInput.Scan() {

		input := userInput.Text()

		for strings.ToLower(input) != "exit" {
			commands.ExecuteCommand(input)

			userInput.Scan()
			input = userInput.Text()
		}

	} else if userInput.Err().Error() != "" {
		fmt.Println("Error while retrieving the command. Please try again")
	}
}
