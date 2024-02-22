package routes

import (
	"encoding/json"
	"hsmyc/gocms/handlers"
	"hsmyc/gocms/models"
	"net/http"

	"github.com/a-h/templ"
)

var singleType models.SingleTypeModel

func SingleTypePost() {
	http.HandleFunc("/api/singletype/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&singleType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handlers.InsertSingleType(singleType)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("SingleType created successfully"))
	})
}

func SingleTypeCreatePage() {
	http.HandleFunc("/cms/singletype", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(handlers.SingleTypeCreatePageHandler()).ServeHTTP(w, r)
	})
}
