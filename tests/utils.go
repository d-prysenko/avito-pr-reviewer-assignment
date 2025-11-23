package tests

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T) *sql.DB {
	cfg := ReadTestConfig()
	db := databaseConnect(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Schema)
	gooseMigrate(t, db)

	return db
}

func databaseConnect(host string, port int, user string, password string, dbname string, schema string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=%s",
		host, port, user, password, dbname, schema)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}

	return db
}

func gooseMigrate(t *testing.T, db *sql.DB) {
	fsys := os.DirFS("../")

	goose.SetBaseFS(fsys)

	if err := goose.SetDialect("postgres"); err != nil {
		require.NoError(t, err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		require.NoError(t, err)
	}

	t.Cleanup(func() {
		err := goose.Reset(db, "migrations")
		require.NoError(t, err)
		db.Close()
	})
}
