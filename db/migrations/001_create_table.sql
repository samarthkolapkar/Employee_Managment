-- Write your migrate up statements here
CREATE TABLE employee (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    middle_name TEXT,
    last_name TEXT NOT NULL,
    pan_number TEXT UNIQUE NOT NULL,
    address_street TEXT NOT NULL,
    address_city TEXT NOT NULL,
    address_state TEXT NOT NULL,
    address_zip TEXT NOT NULL,
    age INT NOT NULL
);

---- create above / drop below ----
SELECT 1;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
