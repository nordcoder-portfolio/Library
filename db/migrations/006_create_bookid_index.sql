-- +goose Up
CREATE INDEX idx_author_book_book_id ON author_book(book_id);

-- +goose Down
DROP INDEX idx_author_book_book_id;