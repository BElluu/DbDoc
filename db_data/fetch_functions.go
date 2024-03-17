package dbdata

import (
	"database/sql"
	"fmt"
)

func FetchFunctions(DB *sql.DB) {

	funtionNames, err := fetchFunctionsName(DB)
	if err != nil {
		fmt.Println("Error getting functions name:", err)
		return
	}
	for _, name := range funtionNames {
		fetchFuntionParameters(DB, name)
	}
	//fetchFuntionParameters(DB, "fn_CommitChanges_Validator")
}

func fetchFunctionsName(DB *sql.DB) ([]string, error) {

	var functionNames []string

	query := `
	SELECT name
	FROM sys.objects
	WHERE type IN ('FN', 'IF') -- FN for scalar-valued functions, IF for inline table-valued functions
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
		var functionName string
		if err := rows.Scan(&functionName); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		fmt.Println("Function Name:", functionName)
		functionNames = append(functionNames, functionName)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return nil, err
	}

	return functionNames, nil
}

func fetchFuntionParameters(DB *sql.DB, functionName string) {
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
            o.type = 'FN'
        AND 
            o.name = '` + functionName + `'
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
		var functionName, dataType string
		var maxLength int
		var isOutput bool
		if err := rows.Scan(&functionName, &dataType, &maxLength, &isOutput); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		var direction string
		if isOutput {
			direction = "OUTPUT"
			functionName = "Return value"
		} else {
			direction = "INPUT"
		}
		fmt.Printf("Function: %s, Data Type: %s, Max Length: %d, Direction: %s\n", functionName, dataType, maxLength, direction)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return
	}
}
