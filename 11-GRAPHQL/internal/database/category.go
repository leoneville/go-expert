package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(ctx context.Context, name, description string) (*Category, error) {
	id := uuid.New().String()
	query := `INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)`

	_, err := c.db.ExecContext(ctx, query, id, name, description)
	if err != nil {
		return nil, err
	}
	return &Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) FindAll(ctx context.Context) ([]Category, error) {
	rows, err := c.db.QueryContext(ctx, "SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}

	return categories, nil
}

func (c *Category) FindByCourseID(ctx context.Context, courseID string) (*Category, error) {
	query := `
			SELECT ca.id, ca.name, ca.description 
			FROM categories ca
			JOIN courses co ON ca.id = co.category_id
			WHERE co.id = $1
	`

	var id, name, description string
	if err := c.db.QueryRowContext(ctx, query, courseID).Scan(&id, &name, &description); err != nil {
		return nil, err
	}

	return &Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) FindByID(ctx context.Context, categoryID string) (*Category, error) {
	query := `SELECT id, name, description FROM categories WHERE id = $1`

	var id, name, description string
	if err := c.db.QueryRowContext(ctx, query, categoryID).Scan(&id, &name, &description); err != nil {
		return nil, err
	}

	return &Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) UpdateCategory(ctx context.Context, categoryID, name, description string) (*Category, error) {
	query := `UPDATE categories SET name = $1, description = $2 WHERE id = $3 RETURNING name, description`

	var dbName, dbDescription string
	if err := c.db.QueryRowContext(
		ctx, query, name, description, categoryID,
	).Scan(&dbName, &dbDescription); err != nil {
		return nil, err
	}
	return &Category{
		ID:          categoryID,
		Name:        dbName,
		Description: dbDescription,
	}, nil
}

func (c *Category) DeleteCategory(ctx context.Context, categoryID string) (*Category, error) {
	deleteCoursesQuery := `DELETE FROM courses WHERE category_id = $1`
	deleteCategoryQuery := `DELETE FROM categories WHERE id = $1 RETURNING name, description`

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, deleteCoursesQuery, categoryID); err != nil {
		return nil, err
	}

	var name, description string
	if err := tx.QueryRowContext(ctx, deleteCategoryQuery, categoryID).Scan(&name, &description); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &Category{
		ID:          categoryID,
		Name:        name,
		Description: description,
	}, nil
}
