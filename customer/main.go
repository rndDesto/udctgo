package main

import "fmt"

type Customer struct {
	ID        int
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

func main() {

	mantul := new(Customer)
	fmt.Println(mantul)
}
