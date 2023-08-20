package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmrashed/rest-api-golang-mongodb/internal/router"
)

func main() {
	r := router.SetupRoutes()

	// Get server address and port
	serverAddr := "localhost"
	serverPort := "8080"

	// Print all routes with complete URL
	fmt.Println("Available routes:")
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tmpl, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		// Print the complete URL
		fmt.Printf("http://%s:%s%s\n", serverAddr, serverPort, tmpl)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Start the server
	serverURL := fmt.Sprintf("%s:%s", serverAddr, serverPort)
	fmt.Printf("Server is running at http://%s\n", serverURL)
	log.Fatal(http.ListenAndServe(serverURL, r))
}
