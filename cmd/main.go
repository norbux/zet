package main

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/norbux/zet/config"
	"github.com/norbux/zet/data"

	err "github.com/norbux/zet/pkg/err"
)

type Zet struct {
	id    string
	title string
	tags  any
}

// TODO: Implement flags (should Charm be considered?)
// TODO: Write README
func main() {
	cfg := config.NewConfig()

	db, e := data.CreateDb(cfg.DatabaseName)
	err.Check(e)

	defer db.Close()

	statement, e := db.Prepare("CREATE TABLE IF NOT EXISTS zet (id varchar(30) PRIMARY KEY, title varchar(256), tags text NULL)")
	if e != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table zet!")
	}
	statement.Exec()

	const timeFormat = "2006-01-02:04-05-06"
	now := time.Now()
	date := now.Format(timeFormat)

	// Insert
	statement, e = db.Prepare("INSERT INTO zet (id, title, tags) VALUES (?, ?, ?)")
	err.Check(e)

	statement.Exec(date, "test", "#tag1 #tag2")
	log.Println("Inserted the zet into database!")
	statement.Close()

	// Query
	rows, e := db.Query("SELECT * FROM zet")
	err.Check(e)
	defer rows.Close()

	for rows.Next() {
		zet := &Zet{}
		e = rows.Scan(&zet.id, &zet.title, &zet.tags)
		err.Check(e)

		log.Printf("id:%s, title:%s, tags:%v", zet.id,
			zet.title, zet.tags)
	}
}
