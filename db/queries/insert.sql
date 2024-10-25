-- name: InsertEmployee :one
INSERT INTO employee_maker (
    employee_data,
    status
) 
VALUES ($1, $2)
RETURNING id;