package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/matthieutran/movie-db/backend/ports"
	"github.com/matthieutran/movie-db/backend/service"
)

const (
	HOST = ""
	PORT = 8080
)

func main() {
	var wg sync.WaitGroup

	// Create Context
	ctx := context.Background()

	// Create application object
	app := service.NewApplication(ctx)

	// Create HTTP Server
	server := ports.NewHttpServer(app)
	// Run HTTP Server in new goroutine
	wg.Add(1)
	go func() {
		err := server.StartHttpServer(wg)(HOST, PORT)
		if err != nil {
			fmt.Println("Could not start HTTP Server:", err)
		}
	}()

	// Block until servers are shut down
	wg.Wait()
}
