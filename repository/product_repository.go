package repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func CreateProduct(product models.Product) (models.Product, error) {
	_, err := db.Conn.NamedExec("INSERT INTO products (name, description, cod, price, ipi, company_id, user_id) VALUES (:name, :description, :cod, :price, :ipi, :company_id, :user_id)", product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func GetProductByID(productID int) (models.Product, error) {
	var product models.Product
	err := db.Conn.Get(&product, "SELECT * FROM products WHERE id = $1", productID)
	if err != nil {
		return product, err
	}
	return product, nil
}
func GetAllProduts() ([]models.Product, error) {
	var products []models.Product
	err := db.Conn.Select(&products, "SELECT * FROM products")
	if err != nil {
		return products, err
	}
	return products, nil
}

func UpdateProduct(product models.Product) (models.Product, error) {

	_, err := db.Conn.NamedExec("UPDATE products SET name = :name, description = :description, cod= :cod , price= :price, ipi= :ipi, company_id = :company_id WHERE id = :id", product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func DeleteProduct(productID int) error {
	_, err := db.Conn.Exec("DELETE FROM products WHERE id = $1", productID)
	if err != nil {
		return err
	}
	return nil
}
