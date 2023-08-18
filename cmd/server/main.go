package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"restapi/api/app"
	"restapi/api/routes"
	"restapi/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Initialize the application
	application := app.NewApplication(cfg)

	// Set up routes
	router := routes.SetupRoutes(application)
    print(fmt.Sprintf(":%d", cfg.PORT))
	// Initialize the server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.PORT),
		Handler: router,
	}

	// Start the server in a separate goroutine
	go func() {
		fmt.Printf("Server started on port %d\n", cfg.PORT)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error:", err)
		}
	}()

	// Handle graceful shutdown
	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	fmt.Println("\nShutting down the server...")

	// Clean up resources, database connections, etc. if needed

	if err := server.Shutdown(nil); err != nil {
		fmt.Println("Error during server shutdown:", err)
	}
	fmt.Println("Server gracefully stopped.")
}
