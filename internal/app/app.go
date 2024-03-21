package app

import (
	"bufio"
	"fmt"
	"github.com/OsipyanG/MiniShell/internal/command"
	"github.com/OsipyanG/MiniShell/internal/process"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	// Инициализация менеджера процессов
	processManager := process.New()

	// Настройка обработчика сигналов для корректного завершения процессов
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		fmt.Println("\nReceived an interrupt, stopping processes...")
		processManager.KillAllProcesses()
		os.Exit(0)
	}()

	// Цикл чтения и выполнения команд
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("$ ")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			fmt.Println("Exiting...")
			break
		} else if line == "" {
			continue
		}
		command.ExecuteCommand(line, processManager)
		fmt.Print("$ ")
	}

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}

	// Завершаем все процессы перед выходом
	processManager.KillAllProcesses()
}
