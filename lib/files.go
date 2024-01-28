package lib

import (
	"errors"
	"os"
)

func GetTodosFile() (file *os.File, err error) {
	todoFile, err := os.OpenFile("todos.json", os.O_RDWR, 0666)
	if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create("todos.json")
		if err == nil {
			todoFile, err = os.OpenFile("todos.json", os.O_RDWR, 0666)
		}
	}

	return todoFile, err
}
