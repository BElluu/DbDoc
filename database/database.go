package database

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func InitDb(server, user, password, database, port string) {
	connectionString := "server=" + server + ";user id=" + user + ";password=" + password + ";port=" + port + ";database=" + database + ";encrypt=disable"

	// Nawiązanie połączenia z bazą danych
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	// Sprawdzenie połączenia
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}

	log.Println("Connected to database successfully")
	DB = db
}
