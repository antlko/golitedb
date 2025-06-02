package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Id       int64  `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return UserRepo{db: db}
}

func (u UserRepo) GetById(ctx context.Context, id int64) (User, error) {
	var user User
	query := "SELECT * FROM dbusers WHERE id = ?"
	if err := u.db.GetContext(ctx, &user, query, id); err != nil {
		return User{}, fmt.Errorf("get user by id: %w", err)
	}
	return user, nil
}

func (u UserRepo) GetByLogin(ctx context.Context, login string) (User, error) {
	var user User
	query := "SELECT * FROM dbusers WHERE login = ?"
	if err := u.db.GetContext(ctx, &user, query, login, login); err != nil {
		return User{}, fmt.Errorf("get user by login: %w", err)
	}
	return user, nil
}

func (u UserRepo) Insert(ctx context.Context, user User) error {
	query := "INSERT INTO dbusers (login, password) VALUES (:login, :email, :password)"
	_, err := u.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}
	return nil
}
