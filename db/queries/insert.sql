-- name: InsertEmployee :one
INSERT INTO employee_maker (
    employee_data,
    status
) 
VALUES ($1, $2)
RETURNING id;

-- name: InsertIntoEmployee :one
INSERT INTO employee(
    first_name,
    middle_name,
    last_name,
    pan_number,
    address_street,
    address_city,
    address_state,
    address_zip,
    age
)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
RETURNING id;
