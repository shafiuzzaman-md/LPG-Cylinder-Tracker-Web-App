package main

import (
	"database/sql"
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

type ScanInfo struct {
	IdScan    string
	Longitude string
	Latitude  string
	UserID    string
	SKU       string
	FillDate  string
	Phone     string
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

		quryScan := "select  idscan, longitude,latitude,user_id, sku, date, phone_identity from scan"

		rows, err := db.Query(quryScan)

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var scans []ScanInfo
		for rows.Next() {
			var temp ScanInfo
			err = rows.Scan(&temp.IdScan, &temp.Longitude, &temp.Latitude, &temp.UserID, &temp.SKU, &temp.FillDate, &temp.Phone)

			scans = append(scans, temp)
			//fmt.Println(temp)
		}

		tmpl.Execute(w, struct{ ScanData []ScanInfo }{scans})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
