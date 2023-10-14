
-- +migrate Up
ALTER TABLE students ADD suspended BOOLEAN NOT NULL DEFAULT FALSE;


-- +migrate Down
ALTER TABLE students DROP COLUMN suspended;
