package main

import (
	"fmt"
	"strconv"
)

type Inventory struct {
	products map[int]*Product
}

func (inv *Inventory) AddProduct(p *Product) {
	if _, ok := inv.products[p.id]; ok {
		fmt.Printf("Product id already exists: %s\n", strconv.Itoa(p.id))
		return
	} else {
		inv.products[p.id] = p
	}
}

func (inv Inventory) Value() (amount int) {
	for _, product := range inv.products {
		amount += product.Value()
	}
	return
}

func (inv *Inventory) Increase(id int, amount int) {
	p, err := inv.GetProductByID(id)
	if err == nil {
		p.Increase(amount)
	} else {
		fmt.Println(err)
	}
}

func (inv *Inventory) Decrease(id int, amount int) {
	p, err := inv.GetProductByID(id)
	if err == nil {
		p.Decrease(amount)
	} else {
		fmt.Println(err)
	}
}

func (inv *Inventory) GetProductByID(id int) (p *Product, err error) {
	if p, ok := inv.products[id]; ok {
		return p, nil
	} else {
		return nil, fmt.Errorf("Product id does not exist: %s", strconv.Itoa(id))
	}
}

func (inv *Inventory) Remove(id int) {
	delete(inv.products, id)
}
