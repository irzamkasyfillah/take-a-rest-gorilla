-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_history (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    action VARCHAR(25) NOT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users_history;
-- +goose StatementEnd
