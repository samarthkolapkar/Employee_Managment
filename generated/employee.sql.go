// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: employee.sql

package generated

import (
	"context"
	"database/sql"
)

const employeeList = `-- name: EmployeeList :many
SELECT 
    e.first_name,
    e.middle_name,
    e.last_name,
    e.pan_number,
    e.address_street,
    e.address_city,
    e.address_state,
    e.address_zip,
    e.age
FROM 
    employee e
WHERE 
    $1::VARCHAR IS NULL 
    OR $1::VARCHAR = '' 
    OR (
        CONCAT_WS(' ',
            first_name,
            middle_name,
            last_name
        ) ILIKE '%' || $1::VARCHAR || '%'
    )
`

type EmployeeListRow struct {
	FirstName     string
	MiddleName    sql.NullString
	LastName      string
	PanNumber     string
	AddressStreet string
	AddressCity   string
	AddressState  string
	AddressZip    string
	Age           int32
}

// db/queries/employee.sql
// db/queries/employee.sql
func (q *Queries) EmployeeList(ctx context.Context, dollar_1 string) ([]EmployeeListRow, error) {
	rows, err := q.db.QueryContext(ctx, employeeList, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EmployeeListRow
	for rows.Next() {
		var i EmployeeListRow
		if err := rows.Scan(
			&i.FirstName,
			&i.MiddleName,
			&i.LastName,
			&i.PanNumber,
			&i.AddressStreet,
			&i.AddressCity,
			&i.AddressState,
			&i.AddressZip,
			&i.Age,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}