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

func (pm *Manager) AddProcess(cmd *exec.Cmd) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	pm.processes = append(pm.processes, cmd)
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error running command: %s\n", err)
	}
}

func (pm *Manager) RemoveProcess(cmd *exec.Cmd) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	for i, process := range pm.processes {
		if process == cmd {
			pm.processes = append(pm.processes[:i], pm.processes[i+1:]...)

			break
		}
	}
}

func (pm *Manager) KillAllProcesses() {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	for _, v := range pm.processes {
		cmd := v
		if cmd.Process != nil {
			err := cmd.Process.Kill()
			if err != nil {
				fmt.Printf("Failed to kill process: %s\n", err)
			} else {
				fmt.Printf("Process killed: %d\n", cmd.Process.Pid)
			}
		}
	}
}
