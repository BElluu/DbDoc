package dbdata

import (
	"database/sql"
	"fmt"
)

type Procedures struct {
	Data []ProcedureData
}

type ProcedureData struct {
	ObjectName string
	Parameters []ProcedureParameters
	Ident      int64
}

type ProcedureParameters struct {
	Name      string
	DataType  string
	MaxLength int64
	IsOutput  bool
}

func FetchProcedures(DB *sql.DB) (ProceduresObjects Procedures) {
	var Objects Procedures
	procedureNames, err := fetchProceudresName(DB)
	if err != nil {
		fmt.Println("Error getting procedures name:", err)
		return
	}

	for _, name := range procedureNames {
		params := fetchProcedureParameters(DB, name)
		Objects.Data = append(Objects.Data, ProcedureData{ObjectName: name, Parameters: params})
	}

	return Objects
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
		procedureNames = append(procedureNames, procedureName)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return procedureNames, nil

}

func fetchProcedureParameters(DB *sql.DB, procedureName string) (params []ProcedureParameters) {
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
		params = append(params, ProcedureParameters{Name: parameterName, DataType: dataType, MaxLength: int64(maxLength), IsOutput: isOutput})
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over results:", err)
		return
	}
	return params
}
