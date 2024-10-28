-- Write your migrate up statements here
ALTER TABLE employee DROP COLUMN address_zip;
alter table employee add COLUMN address_zip BIGINT;
---- create above / drop below ----
SELECT 1;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
