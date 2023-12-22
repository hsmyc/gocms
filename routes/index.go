package routes

import (
	"hsmyc/htmx/handlers"
	"net/http"
)

func Index() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := handlers.IndexHandler()
		component.Render(r.Context(), w)
	})

}
