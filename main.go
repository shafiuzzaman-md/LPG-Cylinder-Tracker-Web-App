package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

/*func apiResponse(w http.ResponseWriter, r *http.Request) {
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
}*/

type info struct {
	Offers      string
	SelectUsers string
	OpenDate    string
	CloseDate   string
	Checked     string
}

/*func SentData() {
	db, err := sql.Open("mysql", "root:hello@tcp(35.200.196.27:3306)/cylindertracker")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO scan VALUES (4, '34.0232','23.5454', 2,'KAJOL','29-06-2020','01773126589' )")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
*/
func main() {

	// Open up database connection.
	db, err := sql.Open("mysql", "root:hello@tcp(35.200.196.27:3306)/cylindertracker")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	//method calling
	//SentData()

	//http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("/static/mydeisgn.css"))))

	//http.HandleFunc("/users", apiResponse)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("home.html"))

		qu := "select longitude, latitude,sku,date,phone_identity from scan"

		rows, err := db.Query(qu)

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var informations []info
		for rows.Next() {
			var temp info
			err = rows.Scan(&temp.Offers, &temp.SelectUsers, &temp.OpenDate, &temp.CloseDate, &temp.Checked)

			informations = append(informations, temp)

		}
		Informations := informations
		//http.Redirect(w, r, "/", 302)
		fmt.Println(Informations)

		tmpl.Execute(w, struct{ Information []info }{informations})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
