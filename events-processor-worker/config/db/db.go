package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"event-processor-worker/utilities"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MYSQL_HOST     string
	MYSQL_DATABASE string
	MYSQL_USER     string
	MYSQL_PASSWORD string
	db             *sql.DB
)

func init() {
	MYSQL_HOST = os.Getenv("MYSQL_HOST")
	MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE")
	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	db = getDbConnection()
}

func getDbConnection() *sql.DB {
	var dbConnectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_DATABASE)

	db, err := sql.Open("mysql", dbConnectionString)
	utilities.ErrorHandler(err, "Failed to connect to the database.")

	return db
}

func CreateTablesIfNotExist() {
	sqlScript, err := os.ReadFile("./config/db/sql-scripts/createTables.sql")
	utilities.ErrorHandler(err, "Failed to read createTables.sql file.")

	statements := strings.Split(string(sqlScript), ";")

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	for _, statement := range statements {
		// Trim whitespace and ignore empty statements
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue
		}

		_, err := db.ExecContext(ctx, statement)
		utilities.ErrorHandler(err, fmt.Sprintf("Failed to execute SQL statement: %s", statement))
	}
	log.Println("Tables created.")
}

func CleanupOnExit() {
	defer db.Close()
}
