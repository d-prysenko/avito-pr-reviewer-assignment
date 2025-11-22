package postgersql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect(host string, port int, user string, password string, dbname string, schema string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=%s",
		host, port, user, password, dbname, schema)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}
