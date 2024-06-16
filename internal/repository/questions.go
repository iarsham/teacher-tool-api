package repository

import (
	"context"
	"database/sql"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"time"
)

type questionsRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) domain.QuestionsRepository {
	return &questionsRepository{
		db: db,
	}
}

func (q *questionsRepository) FindAll() ([]*models.Questions, error) {
	query := `SELECT * FROM questions`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return collectQuestionsRows(rows)
}

func (q *questionsRepository) FindByFile(link string) (*models.Questions, error) {
	query := `SELECT * FROM questions WHERE file = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := q.db.QueryRowContext(ctx, query, link)
	return collectQuestionRow(row)
}

func (q *questionsRepository) FindByID(id uint64) (*models.Questions, error) {
	query := `SELECT * FROM questions WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := q.db.QueryRowContext(ctx, query, id)
	return collectQuestionRow(row)
}

func (q *questionsRepository) Create(question *entities.QuestionRequest, link string, userID uint64) (*models.Questions, error) {
	query := `INSERT INTO questions (lesson, title, grade, level, file, userid) VALUES ($1,$2,$3,$4,$5,$6) RETURNING *`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{question.Lesson, question.Title, question.Grade, question.Level, link, userID}
	row := q.db.QueryRowContext(ctx, query, args...)
	return collectQuestionRow(row)
}

func (q *questionsRepository) Delete(id uint64) error {
	query := `DELETE FROM questions WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return q.db.QueryRowContext(ctx, query, id).Err()
}

func collectQuestionRow(row *sql.Row) (*models.Questions, error) {
	var question models.Questions
	err := row.Scan(
		&question.ID, &question.Lesson,
		&question.Title, &question.Grade,
		&question.Level, &question.Views,
		&question.Used, &question.File,
		&question.UserID, &question.CreatedAt,
	)
	return &question, err
}

func collectQuestionsRows(rows *sql.Rows) ([]*models.Questions, error) {
	var questions []*models.Questions
	for rows.Next() {
		var question models.Questions
		err := rows.Scan(
			&question.ID, &question.Lesson,
			&question.Title, &question.Grade,
			&question.Level, &question.Views,
			&question.Used, &question.File,
			&question.UserID, &question.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		questions = append(questions, &question)
	}
	return questions, nil
}
