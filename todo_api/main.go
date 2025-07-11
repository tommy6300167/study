package main

import (
	"log"
	route "todo_api/router"
)

func main() {
	r := route.SetupRouter()

	log.Println("Starting Todo API server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
