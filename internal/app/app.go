package app

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"
	"syscall"

	"github.com/OsipyanG/MiniShell/internal/command"
	"github.com/OsipyanG/MiniShell/internal/process"
)

func Run() {
	processManager := process.New()
	// defer processManager.KillAllProcesses()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		fmt.Println("\nReceived an interrupt, stopping processes...")
		// processManager.KillAllProcesses()
		os.Exit(0)
	}()

	scanner := bufio.NewScanner(os.Stdin)
	prompt := getPrompt()
	fmt.Print(prompt)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			break
		} else if line == "" {
			continue
		}
		command.ExecuteCommand(line, processManager)
		prompt = getPrompt()
		fmt.Print(prompt)
	}

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
}

func getPrompt() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Failed to get current user:", err)
		os.Exit(1)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Failed to get hostname:", err)
		os.Exit(1)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory:", err)
		os.Exit(1)
	}

	cwd = filepath.Base(cwd)

	// [username@hostname cwd]$
	prompt := fmt.Sprintf("[\033[34m%s\033[0m@\033[32m%s \033[36m%s\033[0m]$ ", currentUser.Username, hostname, cwd)

	return prompt
}
