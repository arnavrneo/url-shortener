package applications

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"net/http"
	"url-shortener/utils"
)

type urlMap struct {
	ShortenedURL string
	OriginalURL  string
	ShortKey     string
}

var urls urlMap
var tpl *template.Template

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	tpl, _ = tpl.ParseGlob("templates/*.html")

	router.Get("/", handleForm)
	router.Post("/shorten", handleShorten)
	router.Get("/short/{id}", handleRedirect)

	// walk the routes
	chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})

	return router
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	// HTML Form
	w.Header().Set("Content-Type", "text/html")
	err := tpl.ExecuteTemplate(w, "form.html", nil)

	if err != nil {
		return
	}
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL is missing", http.StatusBadRequest)
	}

	shortKey := utils.GenerateShortKey()
	shortenedURL := fmt.Sprintf("http://localhost:8000/short/%s", shortKey)

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
	shortKey := chi.URLParam(r, "id")

	if shortKey == "" {
		tpl.ExecuteTemplate(w, "error.html", nil)
		return
	}

	if shortKey == urls.ShortKey {
		originalURL = urls.OriginalURL
	}

	if originalURL == "" {
		tpl.ExecuteTemplate(w, "error.html", nil)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
