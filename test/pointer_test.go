package test

import (
	"fmt"
	"testing"
)

type Customer struct {
	Name string
}
type Binatang struct {
	Name string
}

func UbahNama(customer *Customer) {
	fmt.Println(customer.Name)
	fmt.Println(&customer.Name)
	customer.Name = "Fajri"
	fmt.Println(customer.Name)
	fmt.Println(&customer.Name)
}
func TestPointer(t *testing.T) {
	customer := Customer{
		Name: "Otri",
	}

	//var xx *Customer

	fmt.Println(customer.Name)
	fmt.Println(&customer.Name)
	UbahNama(nil)
	fmt.Println(customer.Name)
	fmt.Println(&customer.Name)

}
