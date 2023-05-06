package helper

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/links")
	if err != nil {
		panic(err.Error())
	}
}

func Insert(longUrl string, shortUrl string) {
	now := time.Now()

	fmt.Println("Inserted at time :", now.Unix())
	later := now.Add(24 * time.Hour)
	expirationTime := later.Unix()
	query := "INSERT INTO links (longLink, shortLink, expirationTime) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		// respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := stmt.Exec(longUrl, shortUrl, expirationTime)
	if err != nil {
		// respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if res == nil {

	}

}

func GetLongUrl(shortUrl string) {
	query := "SELECT * from links where shortLink = ?"
	fmt.Println(shortUrl)
	res, err := db.Query(query, shortUrl)
	if err != nil {
		fmt.Println(err)
		// respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if res == nil {

	}
	var id int
	var longLink string
	var shortLink string
	var exptime int64

	for res.Next() {
		if err := res.Scan(&id, &longLink, &shortLink, &exptime); err != nil {
			fmt.Println(err)
			// handle error
		}
	}

	now := time.Now()
	if exptime < now.Unix(){
      fmt.Println("Link Expired")

	  //DELETE LINK & handle response
	}
	fmt.Println("Table data : ", shortLink, longLink, exptime)
	
	//Return short link and Redirect to it
}
