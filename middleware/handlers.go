package middleware

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres golang driver
)

const (
	host     = "localhost"
	port     = 5432
	user     = "adityapsql"
	password = "123"
	dbname   = "GOREST"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
