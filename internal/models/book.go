package models

import (
	"context"
	"time"
)

type Book struct {
	ID        int       `json:"book_id"`
	Title     string    `json:"book_title"`
	Price     string    `json:"book_price"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func (m *DBModel) GetBooks() ([]*Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var Books []*Book

	query := `SELECT book_id, book_title, book_price, created_at, updated_at 
				FROM books 
				WHERE deleted_at IS NULL 
			`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Book
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Price,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		Books = append(Books, &p)
	}
	return Books, nil
}

func (m *DBModel) GetBook(id int) (Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var b Book

	query := `SELECT book_id, book_title, book_price 
				FROM books 
				WHERE deleted_at IS NULL AND book_id = ?
			`
	if err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&b.ID,
		&b.Title,
		&b.Price,
	); err != nil {
		return b, err
	}

	return b, nil
}

func (m *DBModel) InsertBook(title string, price float64) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO books (book_title, book_price) 
			VALUES (?, ?)`
	result, err := m.DB.ExecContext(ctx, stmt, title, price)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (m *DBModel) DeleteBook(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE books 
			SET deleted_at = UTC_TIMESTAMP()
			WHERE book_id = ?
			`
	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) UpdateBook(id int, title string, price float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE books 
			SET updated_at = UTC_TIMESTAMP(), book_title = ?, book_price = ?
			WHERE book_id = ?
			`
	_, err := m.DB.ExecContext(ctx, stmt,
		title,
		price,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
