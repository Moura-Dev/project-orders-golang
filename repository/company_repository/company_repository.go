package company_repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func Create(company models.Company) (models.Company, error) {

	_, err := db.Conn.NamedExec("INSERT INTO companies (user_id,address_id, phone_id, name, cnpj, fantasy_name, ie, logo) VALUES (:user_id, phone_id, :name, :cnpj, :fantasy_name, :ie, :logo, :address_id);", company)
	if err != nil {
		return company, err
	}
	return company, nil
}

func Update(company models.Company) (models.Company, error) {
	_, err := db.Conn.NamedExec("UPDATE companies SET name = :name, cnpj = :cnpj, fantasy_name = :fantasy_name, ie = :ie, fone = :fone, fone2 = :fone2, logo = :logo, address_id = :address_id WHERE id = :id AND user_id = :user_id", company)
	if err != nil {
		return company, err
	}
	return company, nil
}

func Delete(userID int, companyIDInt int) error {
	_, err := db.Conn.Exec("DELETE FROM companies WHERE user_id = $1 AND id = $2;", userID, companyIDInt)
	if err != nil {
		return err
	}
	return nil
}

func Get(userID int) ([]models.Company, error) {
	var companies []models.Company
	err := db.Conn.Select(&companies, "SELECT * FROM companies WHERE user_id = $1", userID)
	if err != nil {
		return companies, err
	}
	return companies, nil
}
