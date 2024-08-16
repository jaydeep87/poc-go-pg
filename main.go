package main

import (
	// "log"
	"fmt"
	"net/http"
	_ "net/http/pprof" // Import pprof for profiling
	"time"
	"math/rand"
	"sync"

	"github.com/gin-gonic/gin"

	config "github.com/jaydeep87/poc-go-pg/config"
	routes "github.com/jaydeep87/poc-go-pg/routes"
)
// Simulates a CPU-intensive workload
func cpuIntensiveTask() {
	for i := 0; i < 1000000; i++ {
		_ = rand.Float64()
	}
}

// Simulates a memory-intensive workload
func memoryIntensiveTask() {
	data := make([]byte, 0)
	for i := 0; i < 100; i++ {
		data = append(data, byte(i))
		time.Sleep(10 * time.Millisecond) // Simulate some delay
	}
}

func startPprofServer() {
	fmt.Println("Starting pprof server on :6060")
	if err := http.ListenAndServe(":6060", nil); err != nil {
		fmt.Printf("HTTP server error: %v\n", err)
	}
}
func main() {


	// Start the HTTP server for pprof
	go startPprofServer()


	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	go router.Run(":8081")

	var wg sync.WaitGroup

	// Start CPU-intensive tasks
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				cpuIntensiveTask()
			}
		}()
	}

	// Start memory-intensive tasks
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				memoryIntensiveTask()
			}
		}()
	}

	// Wait indefinitely
	wg.Wait()
	// log.Fatal(router.Run(":8081"))
}
