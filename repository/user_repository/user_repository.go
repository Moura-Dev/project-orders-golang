package user_repository

import (
	"context"
	"fmt"
	"github.com/moura-dev/project-orders-golang/db"
	"github.com/moura-dev/project-orders-golang/models"
	"time"
)

func Create(ctx context.Context, user models.User) (models.User, error) {
	writeCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	tx, err := db.Conn.Begin()
	if err != nil {
		fmt.Printf("\nError connecting to the database: %s\n", err)
		return user, err
	}
	defer db.Rollback(tx)

	if err = tx.QueryRowContext(writeCtx, `
			INSERT INTO users (name, email, login, password) 
			VALUES ($1, $2, $3, $4) 
			RETURNING id;`,
		user.Name, user.Email, user.Login, user.Password).
		Scan(&user.Id); err != nil {
		return models.User{}, err
	}

	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Printf("\nError connecting to the database: %v\n", err)
		return user, err
	}

	return user, nil
}

func Get(userId int) (models.User, error) {
	var user models.User

	err := db.Conn.Get(&user, "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1", userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

func Delete(userId int) error {

	tx, err := db.StartTransaction()
	if err != nil {
		return err
	}

	_ = db.Conn.QueryRow("DELETE FROM users WHERE id = $1", userId)

	err = db.CommitTransaction(tx)
	if err != nil {
		err = db.RollbackTransaction(tx)
		if err != nil {
			return err
		}
		return err
	}

	return nil
}

func Login(login string) (models.User, error) {
	var user models.User

	err := db.Conn.Get(&user, "SELECT name, password FROM users WHERE login = $1;", login)
	if err != nil {
		return user, err
	}

	return user, nil
}
