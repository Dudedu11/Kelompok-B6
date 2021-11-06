package main

import (
	"database/sql"
	_ "expvar"
	"fmt"
	"net/http"
	"os"

	"aph-go-service-master/transport"

	"github.com/go-kit/kit/log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "postgres"
	dbpass = "jeje2525"
	dbname = "like_db"
)

func main() {

	initDb()
	defer db.Close()

	logger := log.NewLogfmtLogger(os.Stdout)

	transport.RegisterHttpsServicesAndStartListener()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Log("listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
	}
}

func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	insertStmt := `insert into "feed_like" ("jml_like") values(25)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	insertDynStmt := `insert into "feed_like" ("jml_like") values($1)`
	_, e = db.Exec(insertDynStmt, 30)
	CheckError(e)

	sqlQueryRow() 
	
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}


func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}



func sqlQueryRow() int {
   
    var jml_like  int = 100
    err := db.
        QueryRow("select jumlah_like from feed_like where jumlah_like = ?", jml_like).
        Scan(&jml_like)
    if err != nil {
        panic(err)
    }

    fmt.Printf("jumlah like: %d\n",jml_like)
	return jml_like
}


/*
func SelectLike(jml_like int) int {
	var jumlah_like int 
	sqlStatement := `SELECT jml_like FROM feed_like;`
	row := db.QueryRow(sqlStatement, 3)
	switch err := row.Scan(&jumlah_like); err {
	case sql.ErrNoRows:
  	fmt.Println("No rows were returned!")
	case nil:
  	fmt.Println(jumlah_like)
	default:
  	panic(err)
	}
	return jumlah_like
}
*/


func dbConfig() map[string]string {
	conf := make(map[string]string)

	//will use this later once we set the OS Environment

	/*host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}*/
	conf[dbhost] = "localhost"
	conf[dbport] = "5432"
	conf[dbuser] = "postgres"
	conf[dbpass] = "jeje2525"
	conf[dbname] = "like_db"
	return conf
}
