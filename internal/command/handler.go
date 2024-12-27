package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/OsipyanG/MiniShell/internal/process"
)

func ExecuteCommand(commandLine string, manager *process.Manager) {
	commandParts := strings.Fields(commandLine)
	if len(commandParts) == 0 {
		return
	}

	async := false
	if commandParts[len(commandParts)-1] == "&" {
		async = true
		commandParts = commandParts[:len(commandParts)-1]
	}

	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	if async {
		fmt.Println("Running command in background")
		manager.AddProcess(cmd)
	} else {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}
