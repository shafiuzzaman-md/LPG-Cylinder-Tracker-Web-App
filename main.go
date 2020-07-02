package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// db is the global database connection pool.
var db *sql.DB

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

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Printf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		q := r.URL.Query()

		if len(q) < 3 {
			//log.Println("Url Param 'key' is missing")
			w.Write([]byte("Url Param 'key' is missing"))
			//w.WriteHeader(http.StatusBadRequest)
			SentData("100", "30", "12345")
			return
		} else {
			longitude := strings.Join(q["longitude"], " ")
			//fmt.Println(longitude)
			latitude := strings.Join(q["latitude"], " ")
			sku := strings.Join(q["sku"], " ")
			w.WriteHeader(http.StatusOK)
			result := longitude + latitude + sku
			w.Write([]byte(result))
			//SentData(longitude, latitude, sku)
		}

	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST method requested"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

// initTcpConnectionPool initializes a TCP connection pool for a Cloud SQL
// instance of MySQL.
func initTcpConnectionPool() (*sql.DB, error) {
	// [START cloud_sql_mysql_databasesql_create_tcp]
	var (
		dbUser    = mustGetenv("DB_USER")
		dbPwd     = mustGetenv("DB_PASS")
		dbTcpHost = mustGetenv("DB_TCP_HOST")
		dbName    = mustGetenv("DB_NAME")
	)

	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPwd, dbTcpHost, dbName)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	// [START_EXCLUDE]
	configureConnectionPool(dbPool)
	// [END_EXCLUDE]

	return dbPool, nil
	// [END cloud_sql_mysql_databasesql_create_tcp]
}

// initSocketConnectionPool initializes a Unix socket connection pool for
// a Cloud SQL instance of MySQL.
func initSocketConnectionPool() (*sql.DB, error) {
	// [START cloud_sql_mysql_databasesql_create_socket]
	var (
		dbUser                 = mustGetenv("DB_USER")
		dbPwd                  = mustGetenv("DB_PASS")
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME")
		dbName                 = mustGetenv("DB_NAME")
	)

	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	// [START_EXCLUDE]
	configureConnectionPool(dbPool)
	// [END_EXCLUDE]

	return dbPool, nil
	// [END cloud_sql_mysql_databasesql_create_socket]
}

func configureConnectionPool(dbPool *sql.DB) {
	// [START cloud_sql_mysql_databasesql_limit]

	// Set maximum number of connections in idle connection pool.
	dbPool.SetMaxIdleConns(5)

	// Set maximum number of open connections to the database.
	dbPool.SetMaxOpenConns(7)

	// [END cloud_sql_mysql_databasesql_limit]

	// [START cloud_sql_mysql_databasesql_lifetime]

	// Set Maximum time (in seconds) that a connection can remain open.
	dbPool.SetConnMaxLifetime(1800)

	// [END cloud_sql_mysql_databasesql_lifetime]
}

func SentData(longitude string, latitude string, sku string) {
	var err error

	if os.Getenv("DB_TCP_HOST") != "" {
		db, err = initTcpConnectionPool()
		if err != nil {
			log.Fatalf("initTcpConnectionPool: unable to connect: %s", err)
		}
	} else {
		db, err = initSocketConnectionPool()
		if err != nil {
			log.Fatalf("initSocketConnectionPool: unable to connect: %s", err)
		}
	}
	sql := "INSERT INTO scan VALUES (default," + longitude + "," + latitude + ", 2,'9867411','29-06-2020','01773126589' )"
	//sql := "INSERT INTO scan VALUES (default," + longitude + "," + latitude + ", 2,'KAJOL','29-06-2020','01773126589' )"

	//insert, err := db.Query(sql)
	if _, err = db.Exec(sql); err != nil {
		log.Fatalf("DB.Exec: unable to insert into scan table: %s", err)
	}

	//http.HandleFunc("/", indexHandler)
	/*port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	/*log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}*/

	//db, err := sql.Open("mysql", "root:hello@tcp(35.200.196.27:3306)/cylindertracker")

	//if err != nil {
	//	panic(err.Error())
	//}
	//defer db.Close()
	//sql := "INSERT INTO scan VALUES (default," + longitude + "," + latitude + ", 2,'KAJOL','29-06-2020','01773126589' )"

	//insert, err := db.Query(sql)
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//defer insert.Close()
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
