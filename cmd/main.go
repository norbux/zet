package main

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/norbux/zet/config"
	"github.com/norbux/zet/data"
)

type Zet struct {
	id    string
	title string
	tags  any
}

// TODO: Implement flags (should Charm be considered?)
func main() {
	cfg := config.NewConfig("zet.db")

	db, err := data.CreateDb(cfg.DatabaseName)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS zet (id varchar(30) PRIMARY KEY, title varchar(256), tags text NULL)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table zet!")
	}
	statement.Exec()

	const timeFormat = "2006-01-02:04-05-06"
	now := time.Now()
	date := now.Format(timeFormat)

	// Insert
	statement, err = db.Prepare("INSERT INTO zet (id, title, tags) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec(date, "test", "#tag1 #tag2")
	log.Println("Inserted the zet into database!")
	statement.Close()

	// Query
	rows, err := db.Query("SELECT * FROM zet")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		zet := &Zet{}
		err = rows.Scan(&zet.id, &zet.title, &zet.tags)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("id:%s, title:%s, tags:%v", zet.id,
			zet.title, zet.tags)
	}

}
