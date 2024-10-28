// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: employee.sql

package generated

import (
	"context"
	"database/sql"
	"encoding/json"
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
	AddressZip    sql.NullInt64
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

const getEmployeeDataFrom = `-- name: GetEmployeeDataFrom :one
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
WHERE id=$1
`

type GetEmployeeDataFromRow struct {
	FirstName     string
	MiddleName    sql.NullString
	LastName      string
	PanNumber     string
	AddressStreet string
	AddressCity   string
	AddressState  string
	AddressZip    sql.NullInt64
	Age           int32
}

func (q *Queries) GetEmployeeDataFrom(ctx context.Context, id int32) (GetEmployeeDataFromRow, error) {
	row := q.db.QueryRowContext(ctx, getEmployeeDataFrom, id)
	var i GetEmployeeDataFromRow
	err := row.Scan(
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.PanNumber,
		&i.AddressStreet,
		&i.AddressCity,
		&i.AddressState,
		&i.AddressZip,
		&i.Age,
	)
	return i, err
}

const getEmployeefromMaker = `-- name: GetEmployeefromMaker :many
SELECT
    id,
    employee_data
FROM
    employee_maker
WHERE
    id=$1
    AND status=$2
`

type GetEmployeefromMakerParams struct {
	ID     int32
	Status string
}

type GetEmployeefromMakerRow struct {
	ID           int32
	EmployeeData json.RawMessage
}

func (q *Queries) GetEmployeefromMaker(ctx context.Context, arg GetEmployeefromMakerParams) ([]GetEmployeefromMakerRow, error) {
	rows, err := q.db.QueryContext(ctx, getEmployeefromMaker, arg.ID, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEmployeefromMakerRow
	for rows.Next() {
		var i GetEmployeefromMakerRow
		if err := rows.Scan(&i.ID, &i.EmployeeData); err != nil {
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

const getlistFromEmployeeMaker = `-- name: GetlistFromEmployeeMaker :many
SELECT
    id,
    employee_data,
    status
FROM
    employee_maker
WHERE
    status=$1 
    AND $2::VARCHAR IS NULL 
    OR $2::VARCHAR = '' 
    OR (
        CONCAT_WS(' ',
            id
        ) ILIKE '%' || $2::VARCHAR || '%'
    )
`

type GetlistFromEmployeeMakerParams struct {
	Status  string
	Column2 string
}

func (q *Queries) GetlistFromEmployeeMaker(ctx context.Context, arg GetlistFromEmployeeMakerParams) ([]EmployeeMaker, error) {
	rows, err := q.db.QueryContext(ctx, getlistFromEmployeeMaker, arg.Status, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EmployeeMaker
	for rows.Next() {
		var i EmployeeMaker
		if err := rows.Scan(&i.ID, &i.EmployeeData, &i.Status); err != nil {
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
