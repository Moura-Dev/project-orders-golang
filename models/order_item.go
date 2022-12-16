package models

type OrderItem struct {
	ID        int32  `json:"id"`
	OrderID   int32  `json:"order_id"`
	ProductID int32  `json:"product_id"`
	Quantity  int32  `json:"quantity"`
	Price     int32  `json:"price"`
	CreateAt  string `json:"create_at"`
	UpdateAt  string `json:"update_at"`
}

type OrderItems struct {
	OrderItems []OrderItem `json:"order_items"`
}

type CreateOrderItem struct {
	OrderID   int32 `json:"order_id"`
	ProductID int32 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
	Price     int32 `json:"price"`
}

type UpdateOrderItem struct {
	ID        int32 `json:"id"`
	OrderID   int32 `json:"order_id"`
	ProductID int32 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
	Price     int32 `json:"price"`
}

type DeleteOrderItem struct {
	ID int32 `json:"id"`
}
