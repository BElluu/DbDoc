package dbdata

import (
	"database/sql"
	"fmt"
)

func FetchProcedures(DB *sql.DB) {
	procedureNames, err := fetchProceudresName(DB)
	if err != nil {
		fmt.Println("Error getting procedures name:", err)
		return
	}

	for _, name := range procedureNames {
		fetchProcedureParameters(DB, name)
	}
}

func fetchProceudresName(DB *sql.DB) ([]string, error) {

	var procedureNames []string

	query := `
        SELECT name
        FROM sys.procedures
        ORDER BY name
    `

	// Execute the query
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var procedureName string
		if err := rows.Scan(&procedureName); err != nil {
			return nil, err
		}
		fmt.Println("Procedure Name:", procedureName)
		procedureNames = append(procedureNames, procedureName)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return procedureNames, nil

}

func fetchProcedureParameters(DB *sql.DB, procedureName string) {
	query := `
        SELECT 
            p.name AS parameter_name,
            TYPE_NAME(p.user_type_id) AS data_type,
            p.max_length,
            p.is_output
        FROM 
            sys.parameters p
        INNER JOIN 
            sys.objects o ON p.object_id = o.object_id
        WHERE 
            o.type = 'P'
        AND 
            o.name = '` + procedureName + `'
        ORDER BY 
            p.parameter_id
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
		var parameterName, dataType string
		var maxLength int
		var isOutput bool
		if err := rows.Scan(&parameterName, &dataType, &maxLength, &isOutput); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		var direction string
		if isOutput {
			direction = "OUTPUT"
		} else {
			direction = "INPUT"
		}
		fmt.Printf("Parameter: %s, Data Type: %s, Max Length: %d, Direction: %s\n", parameterName, dataType, maxLength, direction)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return
	}
}
