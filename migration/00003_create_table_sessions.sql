-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions(
    id TEXT NOT NULL,
    hash TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd