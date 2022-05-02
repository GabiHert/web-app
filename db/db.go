package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	fmt.Println("DB CONNECTION STARTED")
	connection := "user=postgres dbname=my_store password=@Ggh18092001 host=localHost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DB CONNECTION OK")
	return db
}
