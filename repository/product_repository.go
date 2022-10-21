package repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func CreateProduct(product models.Product) (models.Product, error) {
	_, err := db.Conn.NamedExec("INSERT INTO products (name, description, price, company_id) VALUES (:name, :description, :price, :company_id)", product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func GetProduct(product models.Product) (models.Product, error) {
	err := db.Conn.Get(&product, "SELECT * FROM products WHERE id = $1", product.ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func UpdateProduct(product models.Product) (models.Product, error) {
	_, err := db.Conn.NamedExec("UPDATE products SET name = :name, description = :description, price = :price, company_id = :company_id WHERE id = :id", product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func DeleteProduct(product models.Product) (models.Product, error) {
	_, err := db.Conn.NamedExec("DELETE FROM products WHERE id = :id", product)
	if err != nil {
		return product, err
	}
	return product, nil
}
