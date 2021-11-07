package database

import (
	"database/sql"
	_ "expvar"
	"fmt"
	"strconv"

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

func InitDb() {
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
	

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	insertStmt := `insert into "feed_like" ("jml_like") values(25)`
	_, e := db.Exec(insertStmt)
	CheckError(e)
/*
	insertDynStmt := `insert into "feed_like" ("jml_like") values($1)`
	_, e = db.Exec(insertDynStmt, 30)
	CheckError(e)
	*/
	//defer db.Close()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

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

func SelectLike() string {

	var jml_lk int
	likeSql := "SELECT jml_like FROM feed_like WHERE jml_like = $1"

	err := db.QueryRow(likeSql,25).Scan(&jml_lk)
	if err != nil {
		if err == sql.ErrNoRows{
			//there no row
		}else{
		panic(err)
		}
	}
	fmt.Printf("jumlah like: %d\n", jml_lk)
	
	return strconv.Itoa(jml_lk)
	
}
