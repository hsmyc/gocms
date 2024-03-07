package routes

import (
	"encoding/json"
	"hsmyc/gocms/handlers"
	"hsmyc/gocms/models"
	"net/http"

	"github.com/a-h/templ"
)

func SingleTypeCreatePage() {
	http.HandleFunc("/cms/singletype", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(handlers.SingleTypeCreatePageHandler()).ServeHTTP(w, r)
	})
}

func CreateSingleType() {
	http.HandleFunc("/api/singletype/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var Doc models.SingleTypeModel
		err := json.NewDecoder(r.Body).Decode(&Doc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		handlers.CreateSingleTypeHandler(Doc)
		response := map[string]string{"status": "success"}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})
}
