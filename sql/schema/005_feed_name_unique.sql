-- +goose Up

ALTER TABLE feeds ADD CONSTRAINT constraint_name UNIQUE (name);

-- +goose Down

ALTER TABLE feeds DROP CONSTRAINT constraint_name;

