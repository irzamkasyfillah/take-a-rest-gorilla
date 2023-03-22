package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/irzam/my-app/api/routes/user"
)

func RegisterRoutes(r *mux.Router) {
	// Index route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("USER API"))
	})

	// Docs routes url
	r.HandleFunc("/api-docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/routes/docs/index.html")
	}).Methods("GET")
	// Docs routes .json
	r.HandleFunc("/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/routes/docs/doc.json")
	}).Methods("GET")

	// User routes
	user.GenerateUserRoutes(r)
}
