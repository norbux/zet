package data

import "database/sql"

func CreateDb(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// TODO: Implement db operations
