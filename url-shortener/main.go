package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
)

type urlMap struct {
	ShortenedURL string
	OriginalURL  string
	ShortKey     string
}

var tpl *template.Template
var urls urlMap

func main() {

	tpl, _ = tpl.ParseGlob("templates/*.html")

	http.HandleFunc("/", handleForm)
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/short/", handleRedirect)

	fmt.Println("URL shortener running on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/shorten", http.StatusSeeOther)
	}

	// HTML Form
	w.Header().Set("Content-Type", "text/html")
	err := tpl.ExecuteTemplate(w, "form.html", nil)

	if err != nil {
		return
	}
}

func handleShorten(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL is missing", http.StatusBadRequest)
	}

	shortKey := generateShortKey()
	shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	urls = urlMap{
		ShortenedURL: shortenedURL,
		OriginalURL:  originalURL,
		ShortKey:     shortKey,
	}

	w.Header().Set("Content-Type", "text/html")
	err := tpl.ExecuteTemplate(w, "urlShorten.html", urls)
	if err != nil {
		return
	}
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	var originalURL string
	shortKey := strings.TrimPrefix(r.URL.Path, "/short/")

	if shortKey == "" {
		http.Error(w, "Shortened key missing", http.StatusBadRequest)
		return
	}
	if shortKey == urls.ShortKey {
		originalURL = urls.OriginalURL
	}
	if originalURL == "" {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
