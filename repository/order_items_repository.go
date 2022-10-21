package repository

import (
	"base-project-api/db"
	"base-project-api/models"
)

func CreateOrderItems(orderItems models.OrderItem) (models.OrderItem, error) {
	_, err := db.Conn.NamedExec("INSERT INTO order_items (order_id, product_id, quantity, total_price) VALUES (:order_id, :product_id, :quantity, :total_price)", orderItems)
	if err != nil {
		return orderItems, err

	}
	return orderItems, nil
}

func DeleteOrderItems(orderItems models.OrderItem) (models.OrderItem, error) {
	_, err := db.Conn.NamedExec("DELETE FROM order_items WHERE id = :id", orderItems)
	if err != nil {
		return orderItems, err
	}
	return orderItems, nil
}

func GetOrderItems(orderItems models.OrderItem) (models.OrderItem, error) {
	err := db.Conn.Get(&orderItems, "SELECT * FROM order_items WHERE order_id = $1", orderItems.OrderID)
	if err != nil {
		return orderItems, err
	}
	return orderItems, nil

}
