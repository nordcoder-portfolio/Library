-- +goose Up
CREATE INDEX idx_book_name ON book(name);

-- +goose Down
DROP INDEX idx_book_name;