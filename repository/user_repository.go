package repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func CreateUser(user models.User) (models.User, error) {

	_, err := db.Conn.NamedExec("INSERT INTO users (first_name, last_name, email, login, password) VALUES (:first_name, :last_name, :email, :login, :password)", user)
	if err != nil {
		return user, err
	}

	return user, nil
}
func GetUserInfo(user models.User) (models.User, error) {

	err := db.Conn.Get(&user, "SELECT * FROM users WHERE id = $1", user.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateSeller(seller models.Seller) (models.Seller, error) {

	_, err := db.Conn.NamedExec("INSERT INTO sellers (user_id, name, tax_id) VALUES (:user_id, :name, :tax_id)", seller)
	if err != nil {
		return seller, err
	}

	return seller, nil
}

func GetSeller(seller models.Seller) (models.Seller, error) {

	err := db.Conn.Get(&seller, "SELECT * FROM sellers WHERE user_id = $1", seller.UserID)
	if err != nil {
		return seller, err
	}

	return seller, nil
}

func DeleteSeller(seller models.Seller) (models.Seller, error) {

	_, err := db.Conn.NamedExec("DELETE FROM sellers WHERE user_id = :user_id", seller.UserID)
	if err != nil {
		return seller, err
	}

	return seller, nil
}
