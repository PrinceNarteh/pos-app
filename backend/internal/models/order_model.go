package models

type Order struct {
	ID           int           `json:"id"`
	Code         string        `json:"code"`
	Total        int64         `json:"total"`
	PPN          int64         `json:"ppn"`
	GrandTotal   int64         `json:"grandTotal"`
	UserID       int           `json:"userId"`
	OrderDetails []OrderDetail `json:"orderDetails"`
}

type OrderDetail struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Qty        int     `json:"qty"`
	TotalPrice float64 `json:"totalPrice"`
	Note       string  `json:"note"`
	ProductID  int     `json:"productId"`
	OrderID    int     `json:"orderId"`
}
