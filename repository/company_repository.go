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
	_, err := db.Conn.NamedExec("UPDATE companies SET name = :name, cnpj = :cnpj, fantasy_name = :fantasy_name, ie = :ie, fone = :fone, fone2 = :fone2, logo = :logo, company_type = :company_type, address_id = :address_id WHERE user_id = :user_id", company)
	if err != nil {
		return company, err
	}
	return company, nil
}

func DeleteCompany(companyIDInt int) error {
	_, err := db.Conn.Exec("DELETE FROM companies WHERE id = $1", companyIDInt)
	if err != nil {
		return err
	}
	return nil
}

func GetCompany(company models.Company) (models.Company, error) {
	err := db.Conn.Get(&company, "SELECT * FROM companies WHERE user_id = $1", company.UserID)
	if err != nil {
		return company, err
	}
	return company, nil
}
