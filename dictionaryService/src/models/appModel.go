package models

import (
	"net/http"
	// _ "github.com/lib/pq"
)

type App struct {
	Router *http.Server
}

// func (app *App) Run(port string) {
// 	log.Println("Server listening on port", port)
// 	log.Fatal(http.ListenAndServe(port, app.Router))
// }

// func respondWithError(w http.ResponseWriter, code int, message string) {
// 	respondWithJSON(w, code, map[string]string{"error": message})
// }

// func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }

// func createRouter(app *App) *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", healthCheck)
// 	http.Handle("/", r)

// 	r.HandleFunc("/api/words", app.getWords).Methods("GET")

// 	return r
// }
