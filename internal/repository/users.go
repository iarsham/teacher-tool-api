package repository

import (
	"context"
	"database/sql"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"time"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) FindAll() ([]*models.Users, error) {
	query := `SELECT * FROM users`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return collectUserRows(rows)
}

func (u *userRepository) FindById(id uint64) (*models.Users, error) {
	query := `SELECT * FROM users WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := u.db.QueryRowContext(ctx, query, id)
	return collectUserRow(row)
}

func (u *userRepository) FindByPhone(phone string) (*models.Users, error) {
	query := `SELECT * FROM users WHERE phone_number = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := u.db.QueryRowContext(ctx, query, phone)
	return collectUserRow(row)
}

func (u *userRepository) Create(user *entities.UserRequest) (*models.Users, error) {
	query := `INSERT INTO users (phone_number, password) VALUES ($1, $2) RETURNING *`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{user.Phone, user.Password}
	row := u.db.QueryRowContext(ctx, query, args...)
	return collectUserRow(row)
}

func (u *userRepository) Update(id uint64, user *entities.UpdateUserRequest) (*models.Users, error) {
	query := `UPDATE users SET phone_number = $1, role=$2 WHERE id = $3 RETURNING *`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{user.Phone, user.Role, id}
	row := u.db.QueryRowContext(ctx, query, args...)
	return collectUserRow(row)
}

func (u *userRepository) UpdatePassword(id uint64, password string) (*models.Users, error) {
	query := `UPDATE users SET password = $1 WHERE id = $2 RETURNING *`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{password, id}
	row := u.db.QueryRowContext(ctx, query, args...)
	return collectUserRow(row)
}

func (u *userRepository) Delete(id uint64) error {
	query := `DELETE FROM users WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return u.db.QueryRowContext(ctx, query, id).Err()
}

func collectUserRow(row *sql.Row) (*models.Users, error) {
	var user models.Users
	err := row.Scan(&user.ID, &user.Phone, &user.Password, &user.CreatedAt, &user.Role)
	return &user, err
}

func collectUserRows(rows *sql.Rows) ([]*models.Users, error) {
	var users []*models.Users
	for rows.Next() {
		var user models.Users
		err := rows.Scan(&user.ID, &user.Phone, &user.Password, &user.CreatedAt, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
