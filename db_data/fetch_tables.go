package dbdata

import (
	"database/sql"
	"fmt"
)

type Tables struct {
	Data []TableData
}

type TableData struct {
	ObjectName string
	Columns    []TableColumns
	Ident      int64
}

type TableColumns struct {
	Name       string
	DataType   string
	MaxLength  int64
	IsNullable bool
}

func FetchTables(DB *sql.DB) (TablesObject Tables) {
	var Objects Tables
	tablesName, err := fetchTablesName(DB)
	if err != nil {
		fmt.Println("Error getting tables name:", err)
		return
	}
	for _, name := range tablesName {
		columns := fetchTableColumns(DB, name)
		Objects.Data = append(Objects.Data, TableData{ObjectName: name, Columns: columns})
	}
	return Objects
}

func fetchTablesName(DB *sql.DB) ([]string, error) {

	var tablesNames []string

	query := `
        SELECT name
        FROM sys.tables
        ORDER BY name
    `

	// Execute the query
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		tablesNames = append(tablesNames, tableName)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return nil, err
	}

	return tablesNames, nil
}

func fetchTableColumns(DB *sql.DB, tableName string) (tables []TableColumns) {

	query := `
	SELECT 
	c.name AS column_name,
	t.name AS data_type,
	c.max_length,
	c.is_nullable
FROM 
	sys.columns c
INNER JOIN 
	sys.types t ON c.user_type_id = t.user_type_id
INNER JOIN 
	sys.tables tb ON c.object_id = tb.object_id
WHERE 
	tb.name = '` + tableName + `'
ORDER BY 
	c.column_id
`

	// Execute the query
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var tableName, dataType string
		var maxLength int
		var isNullable bool
		if err := rows.Scan(&tableName, &dataType, &maxLength, &isNullable); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		tables = append(tables, TableColumns{Name: tableName, DataType: dataType, MaxLength: int64(maxLength), IsNullable: isNullable})
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return
	}
	return tables
}
