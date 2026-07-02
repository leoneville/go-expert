package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(ctx context.Context, name, description, categoryID string) (*Course, error) {
	id := uuid.New().String()
	query := `INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)`

	_, err := c.db.ExecContext(ctx, query, id, name, description, categoryID)
	if err != nil {
		return nil, err
	}

	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

func (c *Course) FindAll(ctx context.Context) ([]Course, error) {
	query := `SELECT id, name, description, category_id FROM courses`
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}

		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}

	return courses, nil
}

func (c *Course) FindByCategoryID(ctx context.Context, categoryID string) ([]*Course, error) {
	query := `SELECT id, name, description, category_id FROM courses WHERE category_id = $1`
	rows, err := c.db.QueryContext(ctx, query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*Course
	for rows.Next() {
		var id, name, description, category_id string
		if err := rows.Scan(&id, &name, &description, &category_id); err != nil {
			return nil, err
		}
		courses = append(courses, &Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryID:  category_id,
		})
	}

	return courses, nil
}

func (c *Course) FindByID(ctx context.Context, courseID string) (*Course, error) {
	query := `SELECT id, name, description, category_id FROM courses WHERE id = $1`

	var id, name, description, category_id string
	if err := c.db.QueryRowContext(ctx, query, courseID).Scan(&id, &name, &description, &category_id); err != nil {
		return nil, err
	}

	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  category_id,
	}, nil
}

func (c *Course) UpdateCourse(ctx context.Context, courseID string, name, description string) (*Course, error) {
	query := `UPDATE courses SET name = $1, description = $2 WHERE id = $3 RETURNING *`

	var id, dbName, dbDescription, category_id string
	if err := c.db.QueryRowContext(
		ctx, query, name, description, courseID,
	).Scan(&id, &dbName, &dbDescription, &category_id); err != nil {
		return nil, err
	}

	return &Course{
		ID:          id,
		Name:        dbName,
		Description: dbDescription,
		CategoryID:  category_id,
	}, nil
}

func (c *Course) DeleteCourse(ctx context.Context, courseID string) (*Course, error) {
	query := `DELETE FROM courses where id = $1 RETURNING *`

	var id, name, description, category_id string
	if err := c.db.QueryRowContext(ctx, query, courseID).Scan(&id, &name, &description, &category_id); err != nil {
		return nil, err
	}

	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  category_id,
	}, nil
}
