package routes

import "net/http"

func Static() {
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend/", fs))
}
