package main

import (
	"devopsowy.pl/dbdoc/database"
	dbdata "devopsowy.pl/dbdoc/db_data"
)

func main() {
	server := "DESKTOP-9REO75F\\DEV"
	port := "1433"
	user := "sa"
	password := "Komenda22!"
	databaseName := "KontrolaWersji"

	database.InitDb(server, user, password, databaseName, port)

	s := dbdata.FetchProcedures(database.DB)
	f := dbdata.FetchFunctions(database.DB)

	println(len(s.Data))
	println(len(f.Data))
	//	fmt.Println("-----------------")
	//	dbdata.FetchFunctions(database.DB)
	//	fmt.Println("--------------")
	//dbdata.FetchTables(database.DB)
}
