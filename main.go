package main

import (
	"fmt"

	cfg "devopsowy.pl/dbdoc/config"
	"devopsowy.pl/dbdoc/database"
	dbdata "devopsowy.pl/dbdoc/db_data"
	tmpl "devopsowy.pl/dbdoc/templates"
)

func main() {
	config, err := cfg.LoadConfig("config.yml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	database.InitDb(config["server"], config["user"], config["password"], config["databaseName"], config["port"])

	s := dbdata.FetchProcedures(database.DB)
	f := dbdata.FetchFunctions(database.DB)
	t := dbdata.FetchTables(database.DB)

	tmpl.FillHTML(s, f, t)

	fmt.Printf("Procedures: %d\n", len(s.Data))
	fmt.Printf("Functions: %d\n", len(f.Data))
	fmt.Printf("Tables: %d\n", len(t.Data))
}
