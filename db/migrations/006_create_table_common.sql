-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS common(
    id SERIAL PRIMARY KEY,
    type TEXT NOT NULL,
    value TEXT NOT NULL,
    is_active BOOLEAN NOT NULL
);
---- create above / drop below ----
SELECT 1;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
