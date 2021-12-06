package service

import (
	"be/config"
	"be/datastruct"
	"log"
)

func MLike(suka datastruct.Menyukai) int64 {

	db := config.CreateConnection()
 
	//defer db.Close()

	sqlStatement := `INSERT INTO menyukai (user_id, feed_id, status_like, tgl_like) VALUES ($1, $2, $3, $4) RETURNING user_id`

	err := db.QueryRow(sqlStatement, suka.User_id, suka.Feed_id, suka.Status_like, suka.Tgl_like).Scan(&suka.User_id)

	if err != nil {
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	}

	return suka.User_id
}

func ShowAllLike(id int64) ([]datastruct.Menyukai, error) {
	db := config.CreateConnection()

	defer db.Close()

	var likes []datastruct.Menyukai


	sqlStatement := `SELECT * FROM menyukai WHERE feed_id=$1 AND status_like IS TRUE `

	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var like datastruct.Menyukai

		err = rows.Scan(&like.User_id, &like.Feed_id, &like.Status_like, &like.Tgl_like)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		likes = append(likes, like)
	}
	return likes, err
}

func ShowLike(id int64) (int64) {
	db := config.CreateConnection()

	defer db.Close()

	var sqlStatement string
	var count int64

	sqlStatement = `SELECT COUNT(*) FROM menyukai WHERE feed_id=$1 AND status_like IS TRUE`


	err := db.QueryRow(sqlStatement, id).Scan(&count)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	return count
}

func UpdateLike(userid int64, feedid int64, suka datastruct.Menyukai) int64 {

	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `UPDATE menyukai SET status_like=$3 WHERE user_id=$1 AND feed_id=$2`

	res, err := db.Exec(sqlStatement, userid, feedid, suka.Status_like)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	return rowsAffected
}