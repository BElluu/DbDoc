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

	//	dbdata.FetchProcedures(database.DB)
	//	fmt.Println("-----------------")
	//	dbdata.FetchFunctions(database.DB)
	//	fmt.Println("--------------")
	dbdata.FetchTables(database.DB)
}
