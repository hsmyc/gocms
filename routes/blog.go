package routes

import (
	"encoding/json"
	"hsmyc/htmx/handlers"
	"hsmyc/htmx/models"
	"net/http"

	"github.com/a-h/templ"
)

var blog models.Blogmodel

func Blog() {
	http.HandleFunc("/blog/", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[6:]
		templ.Handler(handlers.BlogHandler(title)).ServeHTTP(w, r)
	})
}

func CreateBlogPage() {
	http.HandleFunc("/createblogpage", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(handlers.CreateBlogPageHandler()).ServeHTTP(w, r)
	})
}

func CreateBlog() {
	http.HandleFunc("/createblog", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return

		}
		err := json.NewDecoder(r.Body).Decode(&blog)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handlers.CreateBlogHandler(blog)
		w.WriteHeader(http.StatusCreated)
	})
}
