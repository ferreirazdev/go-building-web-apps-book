package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DBHost  = "localhost"
	DBPort  = "5432"
	DBUser  = "postgres"
	DBPass  = "postgres"
	DBDbase = "golang_study"
)

var database *sql.DB

type Page struct {
	Title   string
	Content string
	Date    string
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}

	fmt.Println(pageGUID)

	err := database.QueryRow(`
		SELECT page_title, page_content, page_date from pages WHERE page_guid = $1
		`, pageGUID).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)
	if err != nil {
		log.Println("Couldn't get page: pageGUID")
		log.Println(err.Error())
	}

	html := `<html><head><title>` + thisPage.Title +
		`</title></head><body><h1>` + thisPage.Title + `</h1><div>` +
		thisPage.Content + `</div></body></html>`

	fmt.Fprintln(w, html)
}

func main() {
	dbConn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPass, DBDbase,
	)

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Println("Couldn't connect")
		log.Println(err.Error())
	}

	database = db

	routes := mux.NewRouter()
	routes.HandleFunc("/pages/{guid:[0-9a-zA-Z-]+}", ServePage)
	http.Handle("/", routes)

	http.ListenAndServe(":8080", nil)
}
