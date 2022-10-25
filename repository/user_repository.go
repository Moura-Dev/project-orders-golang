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
func GetUserInfo(userID int) (models.User, error) {
	user := models.User{}
	err := db.Conn.Get(&user, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUser(login string) (models.User, error) {
	user := models.User{}
	err := db.Conn.Get(&user, "SELECT * FROM users WHERE login = $1", login)
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

func GetSellerById(sellerID int) (models.Seller, error) {
	seller := models.Seller{}
	err := db.Conn.Get(&seller, "SELECT id,user_id,name,tax_id FROM sellers WHERE id = $1", sellerID)
	if err != nil {
		return seller, err
	}

	return seller, nil
}

func GetAllSellers() ([]models.Seller, error) {

	var sellers []models.Seller

	err := db.Conn.Select(&sellers, "SELECT id, user_id, name, tax_id FROM sellers")
	if err != nil {
		return sellers, err
	}

	return sellers, nil
}

// delete seller by id
func DeleteSellerByID(sellerID int) error {
	_, err := db.Conn.Exec("DELETE FROM sellers WHERE id = $1", sellerID)
	if err != nil {
		return err
	}

	return nil
}
