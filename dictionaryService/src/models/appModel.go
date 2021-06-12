package models

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type App struct {
	DB     *sql.DB
	Router *http.Server
}

var (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "password"
)

func (app *App) Initialize() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s  sslmode=disable",
		host, port, user, password)

	var err error
	app.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = app.DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB CONNECTED")

	return err
}

func (app *App) CheckDB() bool {
	var (
		err    error
		exists bool
	)

	err = app.DB.QueryRow("SELECT EXISTS(SELECT * FROM information_schema.tables WHERE table_schema = $1 AND table_name = $2);", "public", "words").Scan(&exists)

	if err != nil {
		log.Fatal(err)
	}

	return exists
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
