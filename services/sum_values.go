package services

import (
	"base-project-api/db"
	"github.com/shopspring/decimal"
)

func SumValues(orderID int) (decimal.Decimal, error) {
	var total decimal.Decimal
	err := db.Conn.Get(&total, "SELECT SUM((100-discount) * (price*quantity/100)) AS total FROM order_items WHERE order_id = $1", orderID)
	if err != nil {
		return total, err
	}
	return total, nil
}
