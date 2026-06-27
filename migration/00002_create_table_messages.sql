-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages(
    author TEXT NOT NULL,
    "to" TEXT NOT NULL,
    message TEXT NOT NULL,
    datetime DATETIME
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS messages;
-- +goose StatementEnd