package main

import "fmt"

type Cart struct {
	items map[string]int
}

func NewCart() Cart {
	items := make(map[string]int)
	return Cart{items}
}

func (c *Cart) AddItem(name string, qty int) {
	if val, ok := c.items[name]; ok {
		c.items[name] = val + qty
	} else {
		c.items[name] = qty
	}
}

func (c *Cart) DeleteItem(name string) {
	if _, ok := c.items[name]; ok {
		delete(c.items, name)
	}
}

type CartPrinter struct{}

func (sp CartPrinter) PrintCart(cart Cart, fn func(name string, qty int) string) {
	for name, qty := range cart.items {
		fmt.Println(fn(name, qty))
	}
}

func main() {
	cart := NewCart()
	cart.AddItem("Pisang Hijau", 2)
	cart.AddItem("Semangka Kuning", 3)
	cart.AddItem("Apel Merah", 1)
	cart.AddItem("Apel Merah", 4)
	cart.AddItem("Apel Merah", 2)
	cart.DeleteItem("Semangka Kuning")
	cart.DeleteItem("Semangka Merah")

	cartPrinter := CartPrinter{}
	cartPrinter.PrintCart(cart, func(name string, qty int) string {
		return fmt.Sprintf("%s (%d)", name, qty)
	})
}
