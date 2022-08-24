package process

import (
	"fmt"
	"github.com/amirhnajafiz/procces-monitoring/lock"
	"time"
)

func (p *Process) Run() {
	// Function
	p.Called = 0
	for !p.Terminate {
		// Check for pause
		if p.Pause {
			time.Sleep(3 * time.Second)
			continue
		}
		// Lock
		lock.C.L.Lock()
		lock.Last = p.PID
		// Do
		p.Called++
		p.UpdatedAt = time.Now()
		// Burst
		time.Sleep(time.Second * time.Duration(p.Burst))
		// Unlock
		lock.C.L.Unlock()
		// Waiting
		time.Sleep(time.Second * time.Duration(p.Delay))
	}
}

func (p *Process) Status(i int) string {
	return fmt.Sprintf("%d: Process %d | Task %s | Executed %d | Last Update %s\n", i+1, p.PID, p.Task, p.Called, p.UpdatedAt)
}
