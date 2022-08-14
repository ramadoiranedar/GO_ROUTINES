package go_routines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(i)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	// total CPU
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU", totalCPU)

	// total THREAD
	runtime.GOMAXPROCS(10) // if want to adding THREAD
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	// total Go Routine
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Thread", totalGoroutine)
}
