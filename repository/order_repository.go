package repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func CreateOrder(order models.Order) (models.Order, error) {
	_, err := db.Conn.NamedExec("INSERT INTO orders (status, shipping, user_id, portage_id, customer_id, purchaser, observation) VALUES (:status, :shipping, :user_id, :portage_id, :customer_id, :purchaser, :observation)", order)
	if err != nil {
		return order, err
	}
	return order, nil
}

func GetAllOrders(userID int) ([]models.Order, error) {
	var orders []models.Order
	err := db.Conn.Select(&orders, "SELECT * FROM orders WHERE user_id = $1", userID)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func UpdateOrder(order models.Order) (models.Order, error) {
	_, err := db.Conn.NamedExec("UPDATE orders SET shipping = :shipping, purchaser = :purchaser, observation = :observation WHERE id = :id AND user_id = :user_id", order)
	if err != nil {
		return order, err
	}
	return order, nil
}

func DeleteOrder(orderID int, userID int) error {
	_, err := db.Conn.Exec("DELETE FROM orders WHERE id = $1 AND user_id = $2", orderID, userID)
	if err != nil {
		return err
	}
	return nil
}
