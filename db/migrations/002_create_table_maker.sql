-- Write your migrate up statements here
CREATE TABLE employee_maker (
    id SERIAL PRIMARY KEY,
    employee_data JSONB NOT NULL,
    status TEXT NOT NULL
);
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
