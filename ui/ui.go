package ui

import (
	"bufio"
	"os"
	"todolist/lib"
	"todolist/ui/cli"
	"todolist/ui/tui"
	"todolist/ui/web"
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
		tui.LaunchTUI()
		break
	case WEB:
		web.LaunchWebUI()
		break
	default:
		panic("Invalid UI mode")
	}
}
