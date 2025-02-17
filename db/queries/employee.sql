-- name: EmployeeList :many
-- db/queries/employee.sql
-- db/queries/employee.sql
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
    );

-- name: GetEmployeefromMaker :many
SELECT
    id,
    employee_data
FROM
    employee_maker
WHERE
    id=$1
    AND status=$2;

-- name: GetEmployeeDataFrom :one
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
WHERE id=$1;

-- name: GetlistFromEmployeeMaker :many
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
);


-- name: StatusCount :many
select status::text,COUNT(*)as status_count from employee_maker group by status
UNION ALL
Select status::text ,count(*) as status_count from employee group by status;