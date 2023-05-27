package entity

import (
	"fmt"
	"time"
)

type Order struct {
	ID           int
	ProductId    int
	BuyerEmail   string
	BuyerAddress string
	OrderDate    time.Time
	Product      Product
}

func (o *Order) PrintDetail() {
	fmt.Println("Email Buyer:", o.BuyerEmail)
	fmt.Println("Alamat Buyer:", o.BuyerAddress)
	fmt.Println("Produk : ", o.Product.Name, "Tanggal order : ", o.OrderDate)
}
