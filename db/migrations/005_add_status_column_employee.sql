-- Write your migrate up statements here
ALTER TABLE employee ADD column status VARCHAR(250);
---- create above / drop below ----
SELECT 1;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
