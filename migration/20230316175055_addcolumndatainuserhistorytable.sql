-- +goose Up
-- +goose StatementBegin
ALTER TABLE users_history ADD COLUMN data json NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users_history DROP COLUMN data;
-- +goose StatementEnd
