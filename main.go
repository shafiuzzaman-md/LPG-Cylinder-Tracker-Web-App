package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		q := r.URL.Query()

		if len(q) < 3 {
			//log.Println("Url Param 'key' is missing")
			w.Write([]byte("Url Param 'key' is missing"))
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			longitude := strings.Join(q["longitude"], " ")
			//fmt.Println(longitude)
			latitude := strings.Join(q["latitude"], " ")
			sku := strings.Join(q["sku"], " ")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Saved Successfully"))
			SentData(longitude, latitude, sku)
		}

	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST method requested"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

type UserInfo struct {
	IdUser    string
	UserName  string
	Address   string
	Phone     string
	UserEmail string
	Password  string
	UserType  string
	Date      string
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

func SentData(longitude string, latitude string, sku string) {
	db, err := sql.Open("mysql", "root:hello@tcp(35.200.196.27:3306)/cylindertracker")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	sql := "INSERT INTO scan VALUES (default," + longitude + "," + latitude + ", 2,'KAJOL','29-06-2020','01773126589' )"
	insert, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func main() {

	http.HandleFunc("/scan", apiResponse)
	http.HandleFunc("/", ScanData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ScanData(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("home.html"))
	db, err := sql.Open("mysql", "kajol:kajol123@(192.168.43.140:3306)/cylindertracker")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
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

	quryUser := "select  iduser,name,address,phone,email,password,date,type from user,user_type"

	userRow, err := db.Query(quryUser)
	if err != nil {
		log.Fatal(err)
	}
	defer userRow.Close()
	var UserInformation []UserInfo
	for userRow.Next() {
		var temp UserInfo
		err = userRow.Scan(&temp.IdUser, &temp.UserName, &temp.Address, &temp.Phone, &temp.UserEmail, &temp.Password, &temp.Date, &temp.UserType)
		UserInformation = append(UserInformation, temp)
	}

	tmpl.Execute(w, map[string]interface{}{"ScanData": scans, "UserInfo": UserInformation})

}
