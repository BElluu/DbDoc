package dbdata

import (
	"database/sql"
	"fmt"
)

type Functions struct {
	Data []FunctionData
}

type FunctionData struct {
	ObjectName string
	Parameters []FunctionParameters
}

type FunctionParameters struct {
	Name      string
	DataType  string
	MaxLength int64
	IsOutput  bool
}

func FetchFunctions(DB *sql.DB) (FunctionsObject Functions) {
	var Objects Functions
	funtionNames, err := fetchFunctionsName(DB)
	if err != nil {
		fmt.Println("Error getting functions name:", err)
		return
	}
	for _, name := range funtionNames {
		functions := fetchFuntionParameters(DB, name)
		Objects.Data = append(Objects.Data, FunctionData{ObjectName: name, Parameters: functions})

	}
	return Objects
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
		functionNames = append(functionNames, functionName)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return nil, err
	}

	return functionNames, nil
}

func fetchFuntionParameters(DB *sql.DB, functionName string) (functions []FunctionParameters) {
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
		functions = append(functions, FunctionParameters{Name: functionName, DataType: dataType, MaxLength: int64(maxLength), IsOutput: isOutput})
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return
	}

	return functions
}
