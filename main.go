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
	domain = "http://localhost:3001/"
	// domain = "http://ec2-52-197-102-90.ap-northeast-1.compute.amazonaws.com/"
)

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	fmt.Println("Server successfully connect to database")

	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/{url_id}", Get).Methods("GET")
	router.HandleFunc("/api/v1/urls", Post).Methods("POST")

	fmt.Println("Server successfully setup")
	log.Fatal(http.ListenAndServe(":3001", router))
}

func Get(w http.ResponseWriter, r *http.Request) {
	// 1. record time when GET Method is called
	loc, err := time.LoadLocation("")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Can not load localtion"))
		return
	}
	now := time.Now().In(loc).Format(time.RFC3339)

	// 2. get param from url
	vars := mux.Vars(r)
	url_id := vars["url_id"]

	// 3. using param to search origin url in db
	var originalUrl string
	db := setupDB()
	err = db.QueryRow("SELECT url FROM urls WHERE url_id = $1 AND expireat > $2", url_id, now).Scan(&originalUrl)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("NON-EXISTENT URL or URL EXPIRED"))
		return
	}

    // 4. redirect to the origin website while exist or not expired
	http.Redirect(w, r, originalUrl, http.StatusSeeOther)

	defer db.Close()
}

func Post(w http.ResponseWriter, r *http.Request) {
	// 1. get params from request body
	var req URL
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error in decode")
		return
	}

	if req.Url == "" || req.ExpireAt == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2. insert url, expireDate to db and return url index
	var returnID int
	db := setupDB()
	err = db.QueryRow("INSERT INTO urls(url, expireAt) VALUES($1, $2) returning url_id", req.Url, req.ExpireAt).Scan(&returnID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error in insert")
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)


    // 3. generate JSON file
	id := strconv.Itoa(returnID)
	returnUrl := make(map[string]string)
	returnUrl["id"] = id
	returnUrl["shortUrl"] = domain + id

	// 4. write server response
	response, err := json.Marshal(returnUrl)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(response)

	defer db.Close()
}
