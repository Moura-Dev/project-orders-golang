package repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func CreateAddress(adress models.Address) (models.Address, error) {
	_, err := db.Conn.NamedExec("INSERT INTO adresses (street, number, complement, neighborhood, city, state, zip_code, user_id) VALUES (:street, :number, :complement, :neighborhood, :city, :state, :zip_code, :user_id)", adress)
	if err != nil {
		return adress, err
	}
	return adress, nil
}

func DeleteAddress(adress models.Address) (models.Address, error) {
	_, err := db.Conn.NamedExec("DELETE FROM adresses WHERE user_id = :user_id", adress)
	if err != nil {
		return adress, err
	}
	return adress, nil
}

func GetAddress(adress models.Address) (models.Address, error) {
	err := db.Conn.Get(&adress, "SELECT * FROM adresses WHERE user_id = $1", adress.UserID)
	if err != nil {
		return adress, err
	}
	return adress, nil
}

func UpdateAddress(adress models.Address) (models.Address, error) {
	_, err := db.Conn.NamedExec("UPDATE adresses SET street = :street, number = :number, complement = :complement, neighborhood = :neighborhood, city = :city, state = :state, zip_code = :zip_code WHERE user_id = :user_id", adress)
	if err != nil {
		return adress, err
	}
	return adress, nil
}
