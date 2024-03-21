package command

import (
	"fmt"
	"github.com/OsipyanG/MiniShell/internal/process"
	"os/exec"
	"strings"
)

// ExecuteCommand анализирует входную строку, запускает системные команды синхронно или асинхронно
func ExecuteCommand(commandLine string, manager *process.Manager) {
	commandParts := strings.Fields(commandLine)
	if len(commandParts) == 0 {
		return
	}

	// Определяем, нужно ли выполнять команду асинхронно
	async := false
	if commandParts[len(commandParts)-1] == "&" {
		async = true
		commandParts = commandParts[:len(commandParts)-1] // Убираем "&" из аргументов команды
	}

	if async {
		cmd := exec.Command(commandParts[0], commandParts[1:]...)
		fmt.Println("Running command in background")
		manager.AddProcess(cmd)
	} else {
		cmd := exec.Command(commandParts[0], commandParts[1:]...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
			return
		}
		fmt.Println(string(output))
	}
}
