package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func apiResponse(w http.ResponseWriter, r *http.Request) {
	// Set the return Content-Type as JSON like before
	w.Header().Set("Content-Type", "application/json")

	// Change the response depending on the method being requested
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"name": "Hira"},{"name":"kajol"}]`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST method requested"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

type info struct {
	Offers       string
	Select_Users string
	Open_Date    string
	Close_Date   string
	Checked      string
}

func main() {
	http.HandleFunc("/users", apiResponse)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/temp", func(w http.ResponseWriter, req *http.Request) {

		http.ServeFile(w, req, "Template/temp.html")
	})

	//tmpl := template.Must(template.ParseFiles("../temp.html"))
	db, err := sql.Open("mysql", "root:hello@(35.200.196.27:3306)/cylindertracker")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		qu := "select name, email, phone from user"

		rows, err := db.Query(qu)

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		var informations []info
		for rows.Next() {
			var temp info
			err = rows.Scan(&temp.Offers, &temp.Select_Users, &temp.Open_Date, &temp.Close_Date, &temp.Checked)

			informations = append(informations, temp)

		}
		fmt.Println(informations)

		//tmpl.Execute(w, struct{ Informations []info }{informations})
	})

	log.Fatal(http.ListenAndServe(":8090", nil))
}
