package dbdata

import (
	"database/sql"
	"fmt"
)

type Tables struct {
	Data []TableData
}

type TableData struct {
	ObjectName  string
	Columns     []TableColumns
	Indexes     []TableIndexes
	ForeignKeys []TableForeignKeys
}

type TableColumns struct {
	Name       string
	DataType   string
	MaxLength  int64
	IsNullable bool
}

type TableIndexes struct {
	Name      string
	Columns   string
	Type      string
	IsUnique  bool
	IsPrimary bool
}

type TableForeignKeys struct {
	Name                string
	ReferencedTableName string
	ForeignKeyColumns   string
	ReferencedColumns   string
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
		indexes := fetchTablesIndexes(DB, name)
		foreignKeys := fetchTablesForeignKeys(DB, name)
		Objects.Data = append(Objects.Data, TableData{ObjectName: name, Columns: columns, Indexes: indexes, ForeignKeys: foreignKeys})
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
		var columnName, dataType string
		var maxLength int
		var isNullable bool
		if err := rows.Scan(&columnName, &dataType, &maxLength, &isNullable); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		tables = append(tables, TableColumns{Name: columnName, DataType: dataType, MaxLength: int64(maxLength), IsNullable: isNullable})
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return
	}
	return tables
}

func fetchTablesIndexes(DB *sql.DB, tableName string) (indexes []TableIndexes) {
	query := `
	SELECT
    i.name AS IndexName,
    i.type_desc AS IndexType,
    STRING_AGG(c.name, ', ') AS ColumnNames,
    i.is_unique AS IsUnique,
	i.is_primary_key AS IsPrimary
FROM 
    sys.indexes i
INNER JOIN 
    sys.tables t ON i.object_id = t.object_id
INNER JOIN 
    sys.index_columns ic ON i.object_id = ic.object_id AND i.index_id = ic.index_id
INNER JOIN 
    sys.columns c ON ic.object_id = c.object_id AND ic.column_id = c.column_id
WHERE t.name = '` + tableName + `'
GROUP BY 
    t.name, i.name, i.type_desc, i.is_unique, i.is_primary_key
	`
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var indexName, indexType, columns string
		var isUnique, IsPrimary bool
		if err := rows.Scan(&indexName, &indexType, &columns, &isUnique, &IsPrimary); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		indexes = append(indexes, TableIndexes{Name: indexName, Type: indexType, Columns: columns, IsUnique: isUnique, IsPrimary: IsPrimary})
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return
	}
	return indexes
}

func fetchTablesForeignKeys(DB *sql.DB, tableName string) (foreignKeys []TableForeignKeys) {
	query := `
	SELECT 
    fk.name AS ForeignKeyName,
    OBJECT_NAME(fk.referenced_object_id) AS ReferencedTableName,
    STRING_AGG(COL_NAME(fkc.parent_object_id, fkc.parent_column_id), ', ') AS ForeignKeyColumns,
    STRING_AGG(COL_NAME(fkc.referenced_object_id, fkc.referenced_column_id), ', ') AS ReferencedColumns
FROM 
    sys.foreign_keys fk
INNER JOIN 
    sys.foreign_key_columns fkc ON fk.object_id = fkc.constraint_object_id
WHERE 
    OBJECT_NAME(fk.parent_object_id) = '` + tableName + `'
GROUP BY 
    fk.name, fk.parent_object_id, fk.referenced_object_id
	`

	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var foreignKeyName, refTableName, foreignKeyColumns, refColumns string

		if err := rows.Scan(&foreignKeyName, &refTableName, &foreignKeyColumns, &refColumns); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		foreignKeys = append(foreignKeys, TableForeignKeys{Name: foreignKeyName, ReferencedTableName: refTableName, ForeignKeyColumns: foreignKeyColumns, ReferencedColumns: refColumns})
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return
	}
	return foreignKeys
}
