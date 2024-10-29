-- Write your migrate up statements here
INSERT INTO common (type,value,is_active)
VALUES('DEPARTMENT','HR',true),('DEPARTMENT','ENGINEERING',true),('DEPARTMENT','FINANCE',true),
	('DEPARTMENT','SALES',true),('DEPARTMENT','MARKETING',true),('DEPARTMENT','R&D',true);
---- create above / drop below ----
SELECT  1;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
