package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func getStart(r rune) int {
	return ((int(r) - 32) * 9) + 1
}
func Reader(text string, font string) (string, error) {
	separator := "\n"
	if font == "thinkertoy.txt" {
		separator = "\r\n"
	}
	content, err := os.ReadFile(font)
	if err != nil {
		return "", fmt.Errorf("Error reading the file '%s'", font)
	}

	arts := strings.Split(string(content), separator)
	words := strings.Split(text, "\r\n")
	var result strings.Builder

	for _, w := range words {
		starts := []int{}

		for _, r := range w {
			starts = append(starts, getStart(r))
		}

		if len(starts) == 0 {
			result.WriteString("\n")
			continue
		}

		for i := 0; i < 8; i++ {
			for _, start := range starts {
				result.WriteString(arts[start+i])
			}
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "static/index.html")
}
func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		showErrorAndRedirect(w, "Invalid access. Please use the form.")
		return
	}

	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "Empty input", http.StatusBadRequest)
		return
	}
	font := r.FormValue("font")
	ascii, err := Reader(text, font)
	if err != nil {
		internalServerError(w)

		fmt.Println("ERROR:", err)

		return
	}
	fmt.Fprintf(w, ascii)
}

func showErrorAndRedirect(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, `
		<html>
		<head>
			<meta http-equiv="refresh" content="2;url=/">
		</head>
		<body>
			<h2 style="color:red;">%s</h2>
			<p>Redirecting to home page...</p>
		</body>
		</html>
	`, message)
}

func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, `
		<h1>500 - Internal Server Error</h1>
		<p>Something went wrong on our side.</p>
	`)
}

func main() {

	fmt.Println("http://localhost:8080")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
