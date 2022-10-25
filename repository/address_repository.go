package repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func CreateAddress(address models.Address) (models.Address, error) {
	_, err := db.Conn.NamedExec("INSERT INTO addresses (street, number, zip_code, city, state, district, reference, addressname, user_id) VALUES (:street, :number, :zip_code, :city, :state, :district, :reference, :addressname, :user_id)", address)
	if err != nil {
		return address, err
	}
	return address, nil
}

func DeleteAddress(addressID int) error {
	_, err := db.Conn.Exec("DELETE FROM addresses WHERE id = $1", addressID)
	if err != nil {
		return err
	}
	return nil
}

func GetAddress(userID int) (models.Address, error) {
	var address models.Address
	err := db.Conn.Get(&address, "SELECT * FROM addresses WHERE user_id = $1", userID)
	if err != nil {
		return address, err
	}
	return address, nil
}

func UpdateAddress(address models.Address) (models.Address, error) {
	_, err := db.Conn.NamedExec("UPDATE addresses SET street = :street, number = :number, city = :city, state = :state, district = :district, reference = :reference, addressname=:addressname WHERE id = :id", address)
	if err != nil {
		return address, err
	}
	return address, nil
}
