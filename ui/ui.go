package ui

import (
	"bufio"
	"fmt"
	"os"
	"todolist/lib"
	"todolist/ui/cli"
)

type UIMode string

const (
	CLI UIMode = "cli"
	TUI = "tui"
	WEB = "web"
)

func Start(uiMode UIMode, reader *bufio.Reader, todos *[]*lib.Todo, todoFile *os.File) {
	switch uiMode {
	case CLI:
		cli.LaunchCLI(reader, todos, todoFile)
		break
	case TUI:
		fmt.Println("This mode is unfinished! Come back later!")
		// tui.LaunchTUI()
		break
	case WEB:
		fmt.Println("This mode is unfinished! Come back later!")
		// web.LaunchWebUI()
		break
	default:
		panic("Invalid UI mode")
	}
}
