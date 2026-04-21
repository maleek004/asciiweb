package handlers

import (
	"fmt"
	"net/http"

	"asciiweb/internal/ascii"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ShowErrorAndRedirect(w, "404 - page not found", 404)
		return
	}
	http.ServeFile(w, r, "web/templates/index.html")
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ShowErrorAndRedirect(w, "400 - Invalid access. Please use the form.", 400)
		return
	}

	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "Empty input", http.StatusBadRequest)
		return
	}

	font := r.FormValue("font")
	result, err := ascii.Reader(text, font)
	if err != nil {
		InternalServerError(w, err)
		fmt.Println("ERROR:", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, result)
}

func ShowErrorAndRedirect(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)

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

func InternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, `500 - Internal Server Error`, "\n", err)
}
