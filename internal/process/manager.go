package process

import (
	"fmt"
	"os/exec"
	"sync"
)

type Manager struct {
	mutex     sync.Mutex
	processes []*exec.Cmd
}

func New() *Manager {
	return &Manager{}
}

// AddProcess добавляет процесс в менеджер и запускает его асинхронно
func (pm *Manager) AddProcess(cmd *exec.Cmd) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	pm.processes = append(pm.processes, cmd)
	// Запуск процесса асинхронно
	go func() {
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running command: %s\n", err)
		}
		// Удалить процесс из списка после его завершения
		pm.RemoveProcess(cmd)
	}()
}

func (pm *Manager) RemoveProcess(cmd *exec.Cmd) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	for i, process := range pm.processes {
		if process == cmd {
			// Удаляем процесс из списка, не нарушая порядок остальных элементов
			pm.processes = append(pm.processes[:i], pm.processes[i+1:]...)
			break
		}
	}
}

// KillAllProcesses завершает все запущенные процессы
func (pm *Manager) KillAllProcesses() {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	for i := len(pm.processes) - 1; i >= 0; i-- {
		cmd := pm.processes[i]

		if cmd.Process != nil {
			// Принудительное завершение процесса
			err := cmd.Process.Kill()
			if err != nil {
				// Обработка возможной ошибки при попытке завершить процесс
				fmt.Printf("Failed to kill process: %s\n", err)
			} else {
				fmt.Printf("Process killed: %d\n", cmd.Process.Pid)
			}
		}
	}
	// Очистка списка процессов после их завершения
	pm.processes = []*exec.Cmd{}
}
