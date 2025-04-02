package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/project/library/internal/entity"
)

func NewPostgresRepository(db *pgxpool.Pool, logger *zap.Logger) *postgresRepository {
	return &postgresRepository{
		db:     db,
		logger: logger,
	}
}

var _ BooksRepository = (*postgresRepository)(nil)
var _ AuthorRepository = (*postgresRepository)(nil)

type postgresRepository struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func (r *postgresRepository) rollback(ctx context.Context, tx pgx.Tx) {
	err := tx.Rollback(ctx)
	if err != nil {
		r.logger.Error("Rollback failed", zap.Error(err))
	}
}

func (r *postgresRepository) RegisterAuthor(ctx context.Context, author entity.Author) (entity.Author, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return entity.Author{}, err
	}
	defer r.rollback(ctx, tx)

	const queryAuthor = `
		INSERT INTO author (name) VALUES ($1)
		RETURNING id, created_at, updated_at
	`
	err = tx.QueryRow(ctx, queryAuthor, author.Name).
		Scan(&author.ID, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		return entity.Author{}, err
	}

	const queryAuthorBooks = `
		INSERT INTO author_book (author_id, book_id)
		VALUES ($1, $2)
	`
	for _, bookID := range author.BookIDs {
		_, err := tx.Exec(ctx, queryAuthorBooks, author.ID, bookID)
		if err != nil {
			return entity.Author{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return entity.Author{}, err
	}

	return author, nil
}

func (r *postgresRepository) GetAuthorByID(ctx context.Context, id string) (entity.Author, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return entity.Author{}, err
	}
	defer r.rollback(ctx, tx)

	const query = `
		SELECT a.id, a.name, a.created_at, a.updated_at, ab.author_id
		FROM author a 
		LEFT JOIN author_book ab ON a.id = ab.author_id
		WHERE a.id = $1
	`
	rows, err := tx.Query(ctx, query, id)
	if err != nil {
		return entity.Author{}, err
	}
	defer rows.Close()

	var author entity.Author
	firstRow := true
	for rows.Next() {
		var bookID sql.NullString
		if firstRow {
			if err := rows.Scan(&author.ID, &author.Name, &author.CreatedAt, &author.UpdatedAt, &bookID); err != nil {
				return entity.Author{}, err
			}
			firstRow = false
		} else {
			if err := rows.Scan(new(string), new(string), new(time.Time), new(time.Time), &bookID); err != nil {
				return entity.Author{}, err
			}
		}
		if bookID.Valid {
			author.BookIDs = append(author.BookIDs, bookID.String)
		}
	}
	if err := rows.Err(); err != nil {
		return entity.Author{}, err
	}
	if firstRow {
		return entity.Author{}, entity.ErrBookNotFound
	}

	if err := tx.Commit(ctx); err != nil {
		return entity.Author{}, err
	}
	return author, nil
}

func (r *postgresRepository) ChangeAuthorInfoByID(ctx context.Context, id string, newName string) (entity.Author, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return entity.Author{}, err
	}
	defer r.rollback(ctx, tx)

	const query = `
		UPDATE author
		SET name = $1
		WHERE id = $2
		RETURNING id, name, created_at, updated_at
	`
	var updatedAuthor entity.Author
	err = tx.QueryRow(ctx, query, newName, id).
		Scan(&updatedAuthor.ID, &updatedAuthor.Name, &updatedAuthor.CreatedAt, &updatedAuthor.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Author{}, entity.ErrAuthorNotFound
		}
		return entity.Author{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return entity.Author{}, err
	}
	return updatedAuthor, nil
}

func (r *postgresRepository) GetAuthorBooks(ctx context.Context, authorID string) ([]*entity.Book, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer r.rollback(ctx, tx)

	const queryBooks = `
		SELECT b.id, b.name, b.created_at, b.updated_at
		FROM book b
		JOIN author_book ab ON b.id = ab.book_id
		WHERE ab.author_id = $1
	`
	rows, err := tx.Query(ctx, queryBooks, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*entity.Book
	for rows.Next() {
		var b entity.Book
		if err := rows.Scan(&b.ID, &b.Name, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	for _, b := range books {
		authors, err := r.getBookAuthorsTx(ctx, tx, b.ID)
		if err != nil {
			return nil, err
		}
		b.AuthorsID = authors
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *postgresRepository) getBookAuthorsTx(ctx context.Context, tx pgx.Tx, bookID string) ([]string, error) {
	const query = `
		SELECT a.id
		FROM author a
		JOIN author_book ab ON a.id = ab.author_id
		WHERE ab.book_id = $1
	`
	rows, err := tx.Query(ctx, query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []string
	for rows.Next() {
		var authorID string
		if err := rows.Scan(&authorID); err != nil {
			return nil, err
		}
		authors = append(authors, authorID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *postgresRepository) AddBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return entity.Book{}, err
	}
	defer r.rollback(ctx, tx)

	const queryBook = `
		INSERT INTO book (name) VALUES ($1)
		RETURNING id, created_at, updated_at
	`
	err = tx.QueryRow(ctx, queryBook, book.Name).
		Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return entity.Book{}, err
	}

	if err := r.areAuthorsExistTx(ctx, tx, book.AuthorsID); err != nil {
		return entity.Book{}, entity.ErrAuthorNotFound
	}

	const queryAuthorBook = `
		INSERT INTO author_book (author_id, book_id) VALUES ($1, $2)
	`
	for _, authorID := range book.AuthorsID {
		_, err := tx.Exec(ctx, queryAuthorBook, authorID, book.ID)
		if err != nil {
			return entity.Book{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return entity.Book{}, err
	}
	return book, nil
}

func (r *postgresRepository) GetBookByID(ctx context.Context, bookID string) (entity.Book, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return entity.Book{}, err
	}
	defer r.rollback(ctx, tx)

	const query = `
		SELECT b.id, b.name, b.created_at, b.updated_at, ab.author_id
		FROM book b
		LEFT JOIN author_book ab ON b.id = ab.book_id
		WHERE b.id = $1
	`
	rows, err := tx.Query(ctx, query, bookID)
	if err != nil {
		return entity.Book{}, err
	}
	defer rows.Close()

	var book entity.Book
	firstRow := true
	for rows.Next() {
		var authorID sql.NullString
		if firstRow {
			if err := rows.Scan(&book.ID, &book.Name, &book.CreatedAt, &book.UpdatedAt, &authorID); err != nil {
				return entity.Book{}, err
			}
			firstRow = false
		} else {
			if err := rows.Scan(new(string), new(string), new(time.Time), new(time.Time), &authorID); err != nil {
				return entity.Book{}, err
			}
		}
		if authorID.Valid {
			book.AuthorsID = append(book.AuthorsID, authorID.String)
		}
	}
	if err := rows.Err(); err != nil {
		return entity.Book{}, err
	}
	if firstRow {
		return entity.Book{}, entity.ErrBookNotFound
	}

	if err := tx.Commit(ctx); err != nil {
		return entity.Book{}, err
	}
	return book, nil
}

func (r *postgresRepository) UpdateBookInfoByID(ctx context.Context, bookID, info string, authorIDs []string) (entity.Book, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return entity.Book{}, err
	}
	defer r.rollback(ctx, tx)

	const query = `
		UPDATE book
		SET name = $1
		WHERE id = $2
		RETURNING id, name, created_at, updated_at
	`
	var updatedBook entity.Book
	err = tx.QueryRow(ctx, query, info, bookID).
		Scan(&updatedBook.ID, &updatedBook.Name, &updatedBook.CreatedAt, &updatedBook.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Book{}, entity.ErrAuthorNotFound
		}
		return entity.Book{}, err
	}

	const clearAuthorsQuery = `
		DELETE FROM author_book
		WHERE book_id = $1
	`
	if _, err := tx.Exec(ctx, clearAuthorsQuery, bookID); err != nil {
		return entity.Book{}, err
	}

	if err := r.areAuthorsExistTx(ctx, tx, authorIDs); err != nil {
		return entity.Book{}, entity.ErrAuthorNotFound
	}

	const queryAuthorBook = `
		INSERT INTO author_book (author_id, book_id) VALUES ($1, $2)
	`
	for _, authorID := range authorIDs {
		_, err := tx.Exec(ctx, queryAuthorBook, authorID, updatedBook.ID)
		if err != nil {
			return entity.Book{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return entity.Book{}, err
	}
	return updatedBook, nil
}

func (r *postgresRepository) areAuthorsExistTx(ctx context.Context, tx pgx.Tx, authorIDs []string) error {
	if len(authorIDs) == 0 {
		return nil
	}

	const query = `
		SELECT COUNT(*) FROM author WHERE id = ANY($1)
	`
	var count int
	err := tx.QueryRow(ctx, query, authorIDs).Scan(&count)
	if err != nil {
		return err
	}
	if count != len(authorIDs) {
		return entity.ErrAuthorNotFound
	}
	return nil
}
