package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var listenerPort = ":8090"
var tlsCert = "ssl/cert.pem"
var tlsKey = "ssl/key.pem"

type application struct {
	DB *pgxpool.Pool
}

func (app *application) index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.ServeFile(w, req, "html/index.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		fmt.Fprintf(w, "Post from website! req.PostFrom = %v\n", req.PostForm)
		inputTreeType := req.FormValue("inputTreeType")
		inputTreeLocation := req.FormValue("inputTreeLocation")
		fmt.Fprintf(w, "Type = %s\n", inputTreeType)
		fmt.Fprintf(w, "Location = %s\n", inputTreeLocation)

		sqlStatement := `
		INSERT INTO tree_inventory_v1 (type, location)
		VALUES ($1, $2)`
		_, err := app.DB.Exec(context.Background(), sqlStatement, inputTreeType, inputTreeLocation)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Record inserted to database\n")

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	databaseUrl, isSet := os.LookupEnv("DATABASE_URL")
	if !isSet {
		log.Fatal("env var DATABASE_URL not set")
	}
	conn, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	app := &application{
		DB: conn,
	}

	http.HandleFunc("/", app.index)
	http.HandleFunc("/headers", headers)

	log.Printf("listening on %s\n", listenerPort)
	// log.Fatal(http.ListenAndServe(listenerPort, nil))
	log.Fatal(http.ListenAndServeTLS(listenerPort, tlsCert, tlsKey, nil))
}
