package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type URL struct {
	Url      string `json: "url"`
	ExpireAt string `json: "expireAt"`
}

type ShortenURL struct {
	ID       string `json: "id"`
	ShortUrl string `json: "shortUrl"`
}

const (
	host     = "db"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
	domain = "ec2-52-197-102-90.ap-northeast-1.compute.amazonaws.com"
)

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	fmt.Println("Successfully created connection to database")

	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/{url_id}", get_url).Methods("GET")
	router.HandleFunc("/api/v1/urls", post_url).Methods("POST")

	log.Fatal(http.ListenAndServe(":3001", router))
}

func get_url(w http.ResponseWriter, r *http.Request) {
	loc, err := time.LoadLocation("")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Can not load localtion"))
		return
	}
	now := time.Now().In(loc).Format(time.RFC3339)

	vars := mux.Vars(r)
	url_id := vars["url_id"]

	var originalUrl string
	db := setupDB()
	err = db.QueryRow("SELECT link FROM urls WHERE url_id = $1 AND expireat > $2", url_id, now).Scan(&originalUrl)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("NON-EXISTENT URL or URL EXPIRED"))
		return
	}

	http.Redirect(w, r, originalUrl, http.StatusSeeOther)

	defer db.Close()
}

func post_url(w http.ResponseWriter, r *http.Request) {
	var req URL
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var returnID int
	db := setupDB()
	err = db.QueryRow("INSERT INTO urls(link, expireAt) VALUES($1, $2) returning url_id", req.Url, req.ExpireAt).Scan(&returnID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	id := strconv.Itoa(returnID)
	returnUrl := make(map[string]string)
	returnUrl["id"] = id
	returnUrl["shortUrl"] = domain + id

	response, err := json.Marshal(returnUrl)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(response)

	defer db.Close()
}
