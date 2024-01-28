package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todolist/lib"
	"todolist/ui"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	args := os.Args[1:]

	var mode ui.UIMode
	if len(args) == 0 {
		fmt.Print(`
You must provide an argument for what UI mode to use. Available options are:
--cli
--tui
--web
`)
		return
	}
	switch modeArg := strings.TrimLeft(args[0], "-"); modeArg {
	case "cli":
		mode = ui.CLI
	case "tui":
		mode = ui.TUI
	case "web":
		mode = ui.WEB
	default:
		fmt.Println("Invalid UI mode")
		return
	}

	todoFile, err := lib.GetTodosFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	todos := []*lib.Todo{}
	if loadedTodos, err := lib.LoadTodosJSON(todoFile); err == nil {
		todos = loadedTodos
	}

	ui.Start(mode, reader, &todos, todoFile)
}
