package repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func CreateCompany(company models.Company) (models.Company, error) {
	_, err := db.Conn.NamedExec("INSERT INTO companies (name, tax_id, user_id) VALUES (:name, :tax_id, :user_id)", company)
	if err != nil {
		return company, err
	}
	return company, nil
}

func UpdateCompany(company models.Company) (models.Company, error) {
	_, err := db.Conn.NamedExec("UPDATE companies SET name = :name, tax_id = :tax_id WHERE user_id = :user_id", company)
	if err != nil {
		return company, err
	}
	return company, nil
}

func DeleteCompany(company models.Company) (models.Company, error) {
	_, err := db.Conn.NamedExec("DELETE FROM companies WHERE user_id = :user_id", company)
	if err != nil {
		return company, err
	}
	return company, nil
}

func GetCompany(company models.Company) (models.Company, error) {
	err := db.Conn.Get(&company, "SELECT * FROM companies WHERE user_id = $1", company.UserID)
	if err != nil {
		return company, err
	}
	return company, nil
}
