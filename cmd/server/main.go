package main

import (
	"fmt"
	"net/http"
	"os"

	"asciiweb/internal/handlers"
)

func main() {
	// Use PORT from environment (for Render, Railway, etc.), default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	fmt.Println("Server starting on port", port)
	fmt.Println("http://localhost:" + port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server failed:", err)
		os.Exit(1)
	}
}
