package user_repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/GalahadKingsman/messenger_users/internal/models"
	"strings"
)

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CreateUser(user models.User) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO users (login, first_name, last_name, email, phone) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Login, user.FirstName, user.LastName, user.Email, user.Phone,
	).Scan(&id)
	return id, err
}

func (r *Repo) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (r *Repo) GetUsers(ctx context.Context, filter *models.GetUserFilter) ([]*models.User, error) {
	var (
		query = "SELECT id, login, first_name, last_name, email, phone FROM users"
		args  []interface{}
		where []string
	)

	if filter.Id != nil {
		where = append(where, "id = $1")
		args = append(args, *filter.Id)
	}
	if filter.Login != nil {
		where = append(where, "login = $1")
		args = append(args, *filter.Login)
	}
	if filter.FirstName != nil {
		where = append(where, "first_name = $1")
		args = append(args, *filter.FirstName)
	}
	if filter.LastName != nil {
		where = append(where, "last_name = $1")
		args = append(args, *filter.LastName)
	}
	if filter.Email != nil {
		where = append(where, "email = $1")
		args = append(args, *filter.Email)
	}
	if filter.Phone != nil {
		where = append(where, "phone = $1")
		args = append(args, *filter.Phone)
	}

	// Добавляем WHERE, если есть условия
	if len(where) > 0 {
		query += " WHERE " + strings.Join(where, " AND ")
	}
	if len(where) == 0 {
		return nil, fmt.Errorf("хотя бы одно поле должно быть указано")
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("can not get query context: %w", err)
	}

	users := make([]*models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Login, &user.FirstName, &user.LastName, &user.Email, &user.Phone); err != nil {
			return nil, fmt.Errorf("can not scan row: %w", err)
		}
		users = append(users, &user)
	}

	return users, nil
}
