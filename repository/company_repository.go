package repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func CreateCompany(company models.Company) (models.Company, error) {

	_, err := db.Conn.NamedExec("INSERT INTO companies (user_id, name, cnpj, fantasy_name, ie, fone, fone2, logo, company_type, address_id) VALUES (:user_id, :name, :cnpj, :fantasy_name, :ie, :fone, :fone2, :logo, :company_type, :address_id)", company)
	if err != nil {
		return company, err
	}
	return company, nil
}

func UpdateCompany(company models.Company) (models.Company, error) {
	_, err := db.Conn.NamedExec("UPDATE companies SET name = :name, cnpj = :cnpj, fantasy_name = :fantasy_name, ie = :ie, fone = :fone, fone2 = :fone2, logo = :logo, company_type = :company_type, address_id = :address_id WHERE id = :id AND user_id = :user_id", company)
	if err != nil {
		return company, err
	}
	return company, nil
}

func DeleteCompany(userID int, companyIDInt int) error {
	_, err := db.Conn.Exec("DELETE FROM companies WHERE user_id = $1 AND id = $2", userID, companyIDInt)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCompany(userID int) ([]models.Company, error) {
	var companies []models.Company
	err := db.Conn.Select(&companies, "SELECT * FROM companies WHERE user_id = $1", userID)
	if err != nil {
		return companies, err
	}
	return companies, nil
}
