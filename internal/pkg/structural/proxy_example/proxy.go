package proxy

import (
	"log"
)

type Driven interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	log.Println("I'm driving")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age > 17 {
		c.car.Drive()
		return
	}
	log.Println("Before drive, you must grow little boy")
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}

func Proxy() {
	car := NewCarProxy(&Driver{12})
	car.Drive()
}
