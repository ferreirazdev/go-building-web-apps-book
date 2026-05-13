package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DBHost  = "localhost"
	DBPort  = "5432"
	DBUser  = "postgres"
	DBPass  = "postgres"
	DBDbase = "golang_study"
	PORT    = ":8080"
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

	t, _ := template.ParseFiles("templates/blog.html")

	t.Execute(w, thisPage)
}

func APIPage(w http.ResponseWriter, r *http.Request) {
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

	APIOutput, err := json.Marshal(thisPage)
	fmt.Println(APIOutput)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, thisPage)
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

	routes.HandleFunc("/api/pages", APIPage).Methods("GET").Schemes("https")

	routes.HandleFunc("/api/pages/{guid:[0-9a-zA-Z-]+}", APIPage).Methods("GET").Schemes("https")
	routes.HandleFunc("/page/{guid:[0-9a-zA\\-]+}", ServePage)

	http.Handle("/", routes)

	certificates, err := tls.LoadX509KeyPair("certificate.pem", "key.pem")
	if err != nil {
		log.Fatalf("Couldn't load TLS key pair: %s", err.Error())
	}

	tlsConf := &tls.Config{Certificates: []tls.Certificate{certificates}}

	server := &http.Server{
		Addr:      PORT,
		Handler:   routes,
		TLSConfig: tlsConf,
	}

	log.Printf("Listening on https://localhost%s", PORT)
	log.Fatal(server.ListenAndServeTLS("certificate.pem", "key.pem"))
}
