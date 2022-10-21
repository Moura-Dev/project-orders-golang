package repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func CreateOrder(order models.Order) (models.Order, error) {
	_, err := db.Conn.NamedExec("INSERT INTO order (user_id, product_id, quantity, total_price) VALUES (:user_id, :product_id, :quantity, :total_price)", order)
	if err != nil {
		return order, err
	}
	return order, nil
}

func GetOrders(order models.Order) (models.Order, error) {
	err := db.Conn.Get(&order, "SELECT * FROM order WHERE user_id = $1", order.UserID)
	if err != nil {
		return order, err
	}
	return order, nil
}

func UpdateOrder(order models.Order) (models.Order, error) {
	_, err := db.Conn.NamedExec("UPDATE order SET user_id = :user_id, product_id = :product_id, quantity = :quantity, total_price = :total_price WHERE id = :id", order)
	if err != nil {
		return order, err
	}
	return order, nil
}

func DeleteOrder(order models.Order) (models.Order, error) {
	_, err := db.Conn.NamedExec("DELETE FROM order WHERE id = :id", order)
	if err != nil {
		return order, err
	}
	return order, nil
}
