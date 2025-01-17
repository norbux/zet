package main

import (
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/norbux/zet/config"
	"github.com/norbux/zet/data"

	"github.com/norbux/zet/api"
	"github.com/norbux/zet/pkg/err_check"
	"github.com/norbux/zet/pkg/initialization"
)

// TODO: Move this struct to a "zet" module
type Zet struct {
	id    string
	title string
	tags  any
}

// TODO: Write README
func main() {
	args := os.Args[1:]
	err := initialization.ValidateArgs(args)
	err_check.For(err)

	cfg := config.NewConfig()

	db, err := data.CreateDb(cfg.DatabaseName)
	err_check.For(err)

	defer db.Close()

	api.InitDb(db)
	api.NewRecord(db, "A new Zettel")
	// statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS zet (id varchar(30) PRIMARY KEY, title varchar(256), tags text NULL)")
	// if err != nil {
	// 	log.Println("Error in creating table")
	// } else {
	// 	log.Println("Successfully created table zet!")
	// }
	// statement.Exec()

	// Query
	rows, err := db.Query("SELECT * FROM zet")
	err_check.For(err)
	defer rows.Close()

	for rows.Next() {
		zet := &Zet{}
		err = rows.Scan(&zet.id, &zet.title, &zet.tags)
		err_check.For(err)

		log.Printf("id:%s, title:%s, tags:%v", zet.id,
			zet.title, zet.tags)
	}
}
