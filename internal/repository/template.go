package repository

import (
	"context"
	"database/sql"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"time"
)

type templateRepository struct {
	db *sql.DB
}

func NewTemplateRepository(db *sql.DB) domain.TemplateRepository {
	return &templateRepository{
		db: db,
	}
}

func (t *templateRepository) FindAll() ([]*models.Templates, error) {
	query := `SELECT * FROM template`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := t.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return collectTemplateRows(rows)
}

func (t *templateRepository) FindByFile(link string) (*models.Templates, error) {
	query := `SELECT * FROM template WHERE file = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := t.db.QueryRowContext(ctx, query, link)
	return collectTemplateRow(row)
}

func (t *templateRepository) Create(template *entities.TemplateRequest, link string) (*models.Templates, error) {
	query := `INSERT INTO template (userid, file) VALUES ($1,$2) RETURNING *`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{template.UserID, link}
	row := t.db.QueryRowContext(ctx, query, args...)
	return collectTemplateRow(row)
}

func (t *templateRepository) Delete(templateID uint64) error {
	query := `DELETE FROM template WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return t.db.QueryRowContext(ctx, query, templateID).Err()
}

func collectTemplateRow(row *sql.Row) (*models.Templates, error) {
	var tmpl models.Templates
	err := row.Scan(&tmpl.ID, &tmpl.UserID, &tmpl.File, &tmpl.CreatedAt)
	return &tmpl, err
}

func collectTemplateRows(rows *sql.Rows) ([]*models.Templates, error) {
	defer rows.Close()
	var templates []*models.Templates
	for rows.Next() {
		var template models.Templates
		if err := rows.Scan(&template.ID, &template.UserID, &template.File, &template.CreatedAt); err != nil {
			return nil, err
		}
		templates = append(templates, &template)
	}
	return templates, nil
}
