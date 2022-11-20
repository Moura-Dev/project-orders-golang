package address_repository

import (
	"github.com/moura-dev/project-orders-golang/db"
	"github.com/moura-dev/project-orders-golang/models"
)

func Create(address models.Address) (models.Address, error) {
	_, err := db.Conn.NamedExec("INSERT INTO addresses (street, number, zip_code, city, state, district, reference, addressname, user_id) VALUES (:street, :number, :zip_code, :city, :state, :district, :reference, :addressname, :user_id)", address)
	if err != nil {
		return address, err
	}
	return address, nil
}

func Delete(userID int, addressID int) error {
	_, err := db.Conn.Exec("DELETE FROM addresses WHERE id = $1 AND user_id = $2", addressID, userID)
	if err != nil {
		return err
	}
	return nil
}

func Get(userID int) ([]models.Address, error) {
	var addresses []models.Address
	err := db.Conn.Select(&addresses, "SELECT * FROM addresses WHERE user_id = $1", userID)
	if err != nil {
		return addresses, err
	}
	return addresses, nil
}

func Update(address models.Address) (models.Address, error) {
	_, err := db.Conn.NamedExec("UPDATE addresses SET street = :street, number = :number, city = :city, state = :state, district = :district, reference = :reference, addressname=:addressname WHERE id = :id", address)
	if err != nil {
		return address, err
	}
	return address, nil
}
