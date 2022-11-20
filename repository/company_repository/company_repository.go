package company_repository

import (
	"github.com/moura-dev/project-orders-golang/db"
	"github.com/moura-dev/project-orders-golang/models"
)

func Create(userId int, company models.Company) (models.Company, error) {
	_, err := db.Conn.NamedExec("INSERT INTO companies (created_at, updated_at) VALUES (NOW(), NOW());", company)
	if err != nil {
		return company, err
	}

	_, err = db.Conn.Exec("INSERT INTO user_companies (company_id, user_id) VALUES ($1, $2);", company.Id, userId)
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

func Delete(userID int, companyId int) error {
	_, err := db.Conn.Exec("DELETE FROM companies WHERE id = $1", companyId)
	if err != nil {
		return err
	}
	return nil
}

func Get(userID int) ([]models.Company, error) {
	var companies []models.Company
	err := db.Conn.Select(&companies, "SELECT id FROM companies c JOIN user_companies uc on c.id = uc.company_id WHERE user_id = $1", userID)
	if err != nil {
		return companies, err
	}
	return companies, nil
}
