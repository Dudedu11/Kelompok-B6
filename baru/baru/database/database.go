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
	dbname = "db_like"
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

	//InsertLike()
	
	//defer db.Close()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func dbConfig() map[string]string {
	conf := make(map[string]string)


	conf[dbhost] = "localhost"
	conf[dbport] = "5432"
	conf[dbuser] = "postgres"
	conf[dbpass] = "jeje2525"
	conf[dbname] = "db_like"
	return conf
}

func InsertLike(){
	insertStmt := `insert into "menyukai" ("user_id","feed_id","status_like","tgl_like") values(3,2,TRUE,'2021-11-22')`
	_, e := db.Exec(insertStmt)
	CheckError(e)
}


func SelectLike() string {

	var jml_lk int
	var feed_lk int 

	//likeSql := "SELECT jml_like FROM feed_like WHERE jml_like = $1"
	likeSql := "SELECT feed_id, COUNT(user_id) AS jml_like FROM menyukai WHERE status_like IS TRUE AND feed_id = $1 GROUP BY feed_id"
	err := db.QueryRow(likeSql,1).Scan(&feed_lk, &jml_lk)
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
