package dbdata

import (
	"database/sql"
	"fmt"
)

type Views struct {
	Data []ViewData
}

type ViewData struct {
	ObjectName string
	Columns    []ViewColumns
}

type ViewColumns struct {
	Name     string
	DataType string
}

func FetchViews(DB *sql.DB) (ViewsObject Views) {
	var Objects Views

	viewsName, err := fetchViewsName(DB)
	if err != nil {
		fmt.Println("Error getting views name:", err)
		return
	}
	for _, name := range viewsName {
		columns := fetchViewColumns(DB, name)
		Objects.Data = append(Objects.Data, ViewData{ObjectName: name, Columns: columns})
	}
	return Objects
}

func fetchViewsName(DB *sql.DB) ([]string, error) {
	var viewsNames []string

	query := `
	SELECT name
    FROM sys.views
	ORDER BY name
	`

	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var viewName string
		if err := rows.Scan(&viewName); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		viewsNames = append(viewsNames, viewName)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return nil, err
	}

	return viewsNames, nil
}

func fetchViewColumns(DB *sql.DB, tableName string) (views []ViewColumns) {

	query := `
	SELECT 
    c.name AS column_name,
    t.name AS data_type
	FROM sys.columns c
	INNER JOIN sys.views v ON c.object_id = v.object_id
	INNER JOIN sys.types t ON c.user_type_id = t.user_type_id
	WHERE v.name = '` + tableName + `'
	`

	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var columnName, dataType string
		if err := rows.Scan(&columnName, &dataType); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		views = append(views, ViewColumns{Name: columnName, DataType: dataType})
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return
	}
	return views
}
