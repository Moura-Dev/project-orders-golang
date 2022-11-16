package user_repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func Create(user models.User) (models.User, error) {

	_, err := db.Conn.NamedExec("INSERT INTO users (first_name, last_name, email, login, password) VALUES (:first_name, :last_name, :email, :login, :password)", user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func Get(userId int) (user models.User, err error) {

	err = db.Conn.Get(&user, "SELECT * FROM users WHERE id = $1", userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

func Delete(userId int) error {
	_, err := db.Conn.Exec("DELETE FROM sellers WHERE id = $1", userId)
	if err != nil {
		return err
	}

	return nil
}

func Login(login string) (models.User, error) {
	var user models.User

	err := db.Conn.Get(&user, "SELECT * FROM users WHERE login = $1", login)
	if err != nil {
		return user, err
	}

	return user, nil
}
