package models


import "time"

type Category struct {

	ID int64
	Name string
	Description string 

}

type Product struct {

	ID int64
	CategoryID int64
	Name string
	Price float64
	Stock int
	LowStockThreshold int
	CreatedAt time.Time

}

type Order struct {
	ID  int64
	Status string
	Total float64
	CreatedAt time.Time

}


type OrderItem struct {

	ID int64
	OrderID int64
	ProductID int64
	Quantity int
	UnitPrice float64
	
}