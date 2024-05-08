package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/OsipyanG/MiniShell/internal/process"
)

// ExecuteCommand анализирует входную строку, запускает системные команды синхронно или асинхронно
func ExecuteCommand(commandLine string, manager *process.Manager) {
	commandParts := strings.Fields(commandLine)
	if len(commandParts) == 0 {
		return
	}

	// Определяем, нужно ли выполнять команду асинхронно
	async := false
	interMode := false
	if commandParts[len(commandParts)-1] == "&" {
		async = true
		commandParts = commandParts[:len(commandParts)-1] // Убираем "&" из аргументов команды
	} else if commandParts[len(commandParts)-1] == "&!" {
		commandParts = commandParts[:len(commandParts)-1]
		interMode = true
	}

	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	if async {
		fmt.Println("Running command in background")
		manager.AddProcess(cmd)

	} else if commandParts[0] == "vim" || interMode {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Запускаем команду
		err := cmd.Run()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
		}

	} else {
		output, err := cmd.CombinedOutput()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error executing command: %s\n", err)
			return
		}
		fmt.Println(string(output))

	}
}
