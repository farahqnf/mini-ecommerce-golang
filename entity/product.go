package entity

import "fmt"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
}

func (p *Product) PrintDetail() {
	fmt.Println("ID:", p.ID)
	fmt.Println(p.Name)
	fmt.Println("Description : ", p.Description)
	fmt.Println("Price : ", p.Price)
}
