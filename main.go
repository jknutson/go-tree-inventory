package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

var listenerPort = ":8090"
var tlsCert = "ssl/cert.pem"
var tlsKey = "ssl/key.pem"

type application struct {
	DB *pgxpool.Pool
}

type PageData struct {
	HasFlashMessage   bool
	FlashMessageText  string
	FlashMessageClass string
}

func (app *application) index(w http.ResponseWriter, req *http.Request) {

	tmpl := template.Must(template.ParseFiles("index.html.tmpl"))

	switch req.Method {
	case "GET":
		pageData := PageData{
			HasFlashMessage: false,
		}
		err := tmpl.Execute(w, pageData)
		if err != nil {
			log.Fatal(err)
		}
		// http.ServeFile(w, req, "html/index.html")
	case "POST":
		var pageErrors []string
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if parseFormErr := req.ParseForm(); parseFormErr != nil {
			pageErrors = append(pageErrors, fmt.Sprintf("ParseForm() err: %v", parseFormErr))
		}

		// TODO: render template with flash
		inputTreeType := req.FormValue("inputTreeType")
		inputTreeLocation := req.FormValue("inputTreeLocation")

		// TODO: update sql to insert geom/make_point
		sqlStatement := `
INSERT INTO tree_inventory_v1 (type, location)
		VALUES ($1, $2)`
		_, sqlErr := app.DB.Exec(context.Background(), sqlStatement, inputTreeType, inputTreeLocation)
		if sqlErr != nil {
			pageErrors = append(pageErrors, fmt.Sprintf("SQL insert err: %v", sqlErr))
		}
		pageData := PageData{
			HasFlashMessage:   true,
			FlashMessageClass: "alert-primary",
		}
		if len(pageErrors) > 0 {
			pageData.FlashMessageClass = "alert-danger"
			pageData.FlashMessageText = strings.Join(pageErrors, "\n")
		} else {
			pageData.FlashMessageText = "Record inserted!"
		}
		if tmplErr := tmpl.Execute(w, pageData); tmplErr != nil {
			fmt.Fprintf(w, "template err: %v", tmplErr)
		}

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
