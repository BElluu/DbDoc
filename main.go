package main

import (
	"devopsowy.pl/dbdoc/database"
	dbdata "devopsowy.pl/dbdoc/db_data"
)

func main() {
	server := "localhost"
	port := "1433"
	user := "sa"
	password := "test"
	databaseName := "test"

	database.InitDb(server, user, password, databaseName, port)

	s := dbdata.FetchProcedures(database.DB)
	f := dbdata.FetchFunctions(database.DB)
	t := dbdata.FetchTables(database.DB)

	println(len(s.Data))
	println(len(f.Data))
	println(len(t.Data))
}
