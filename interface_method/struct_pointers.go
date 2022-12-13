package main

import "fmt"

type car struct {
	brand string

	price int
}

func changeCar(c car, brand string, price int) {
	c.price = price
	c.brand = brand
}
func (c *car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	(*c).price = newPrice
}
func main() {
	myCar := car{brand: "Audi", price: 40000}
	anotherCar := car{brand: "Renault", price: 25000}
	fmt.Printf("another car: %#v\n", anotherCar)
	anotherCar.changeCar1("KIA", 30000)
	changeCar(myCar, "KIA", 35000)
	fmt.Printf("another car: %#v \n", anotherCar)
	fmt.Printf("my car: %#v \n", myCar)

}
