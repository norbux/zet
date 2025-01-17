package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// InitDb bootstraps database for fist use
func InitDb(db *sql.DB) error {
	// TODO: Implement necessary validations to avoid re-initialising
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS zet (id varchar(30) PRIMARY KEY, title varchar(256), tags text NULL)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table zet!")
	}
	statement.Exec()

	return nil
}

// NewRecord creates a new zet entry in the database, creates the corresponding file and loads it in the editor
func NewRecord(db *sql.DB, title string) error {
	// TODO: Handle tags input
	const timeFormat = "20060102-030405"
	now := time.Now()
	date := now.Format(timeFormat)

	statement, err := db.Prepare("INSERT INTO zet (id, title, tags) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	statement.Exec(date, title, "#tag1 #tag2")
	statement.Close()

	// TODO: Hande path in configuration
	cmd := exec.Command("touch", date)

	_, err = cmd.Output()
	if err != nil {
		return err
	}

	cmd = exec.Command("vim", date)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error starting Vim:", err)
		return err
	}

	return nil
}
