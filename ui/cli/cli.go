package cli

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"todolist/lib"
)

func getUserInput(str *string, reader *bufio.Reader) {
	input, err := reader.ReadString('\n')
	for err != nil {
		input, err = reader.ReadString('\n')
	}
	*str = strings.Trim(input, " 	\n")
}

func LaunchCLI(reader *bufio.Reader, todos *[]*lib.Todo, todoFile *os.File) {
	fmt.Println("Launching cli...")
	fmt.Println("Available commands are: list, add, delete")
	for {
		fmt.Println("")
		fmt.Println("What would you like to do?")
			
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Something went wrong")	
			break
		}
		fmt.Println("")

		switch strings.Trim(command, " \n") {
		case "list":
			for index, todo := range *todos {
				fmt.Printf(
					"Index: %v	Completed: %v	Created on: %v\nTitle: %v\n%v\n%v\n\n",
					index, todo.Completed(), todo.CreatedAt(), todo.Title(), strings.Repeat("=", 64), todo.Content(),
				)
			}
			break
		case "add":
			var title string
			var content string
			fmt.Println("Enter a title: ")
			getUserInput(&title, reader)
			fmt.Println("Enter content: ")
			getUserInput(&content, reader)

			*todos = append(*todos, lib.NewTodo(title, content))
			lib.StoreTodosJSON(todoFile, todos)
			fmt.Println("You've added a new todo item!")
			break
		case "delete":
			fmt.Println("Enter the index of the item you'd like to delete: ")
			var input string
			getUserInput(&input, reader)
			index, err := strconv.Atoi(input)
			
			if err != nil {
				fmt.Println("That's not an integer!")
				break
			}
			
			if index >= len(*todos) || index < 0  {
				fmt.Println("You've provided an invalid index!")
				break
			}

			*todos = slices.Delete(*todos, int(index), int(index+1))
			fmt.Println("The todo has been deleted!")
			lib.StoreTodosJSON(todoFile, todos)
			break
		default:
			fmt.Println("Invalid command! Available commands are: list, add, delete")
		}
	}
}
