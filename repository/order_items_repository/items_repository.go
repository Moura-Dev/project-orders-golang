package order_items_repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func Create(item models.OrderItem) (models.OrderItem, error) {
	_, err := db.Conn.NamedExec("INSERT INTO order_items (order_id, product_id, quantity, price, discount) VALUES (:order_id, :product_id, :quantity, :price, :discount)", item)
	if err != nil {
		return models.OrderItem{}, err
	}

	return item, nil
}

func Delete(productID int, orderID int) error {
	_, err := db.Conn.Exec("DELETE FROM order_items WHERE product_id = $1 AND order_id = $2", productID, orderID)
	if err != nil {
		return err
	}
	return nil
}

func Get(orderID int) ([]models.OrderItem, error) {
	var items []models.OrderItem
	err := db.Conn.Select(&items, "SELECT order_id,product_id,quantity,price, discount, (100-discount) * (price*quantity/100) AS total FROM order_items WHERE order_id = $1", orderID)

	if err != nil {
		return items, err
	}
	return items, nil
}

func Update(item models.OrderItem) (models.OrderItem, error) {
	query := `
		UPDATE order_items
		SET product_id = :product_id,
			quantity   = :quantity,
			price      = :price,
			discount   = :discount
		WHERE product_id = :product_id;
		`
	_, err := db.Conn.NamedExec(query, item)
	if err != nil {
		return item, err
	}
	return item, nil
}
