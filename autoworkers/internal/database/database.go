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
	fmt.Println("after open", err)
	if err != nil {
		fmt.Println(err)
	}
	err =db.Ping()
	if err!=nil{
		 fmt.Println(err)
	}else{

	}
	return &Database{
	DB: db,
	}
}