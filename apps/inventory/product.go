package main

type Product struct {
	id       int
	price    int
	quantity int
}

func (p Product) Value() (amount int) {
	amount = p.price * p.quantity
	return
}

func (p *Product) Increase(amount int) {
	p.quantity += amount
}

func (p *Product) Decrease(amount int) {
	p.quantity -= amount
	if p.quantity < 0 {
		p.quantity = 0
	}
}
