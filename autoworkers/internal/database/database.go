package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Database struct {
	DB *sql.DB

}

func Constructor() *Database{
	db, err := sql.Open("sqlite", "./taskforge.db")
	query := `
	CREATE TABLE Jobs(
	id TEXT,
	type TEXT,
	payload TEXT,
	status INTEGER,
	result TEXT
	)
	`
	fmt.Println("after open", err)
	if err != nil {
		fmt.Println(err)
	}
	err =db.Ping()
	_, err = db.Exec(query)
	if err!=nil{
		 fmt.Println(err)
	}else{

	}
	return &Database{
	DB: db,
	}
}