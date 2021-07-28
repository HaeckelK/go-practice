// Product Inventory Project
// Create an application which manages an inventory of products.
// Create a product class which has a price, id, and quantity on hand.
// Then create an inventory class which keeps track of various products
// and can sum up the inventory value.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Inventory Manager")
	inv := Inventory{products: make(map[int]*Product)}
	p1 := Product{id: 1, price: 100, quantity: 5}
	inv.AddProduct(&p1)
	p2 := Product{id: 2, price: 25, quantity: 16}
	inv.AddProduct(&p2)

	inv.AddProduct(&p2)

	inv.Increase(2, 4)
	inv.Decrease(1, 4)

	inv.Increase(99, 100)
	inv.Remove(1)
	inv.Remove(99)

	for _, product := range inv.products {
		fmt.Printf("%+v\n", *product)
	}

	fmt.Println("Inventory Value:", strconv.Itoa(inv.Value()))
}
