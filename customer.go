package main

import "errors"

var (
	errOutOfStock = errors.New("requested CD is out of stock")
)

type Customer struct{}

func (c *Customer) Buy(cd *CD) error {
	if cd.Stock == 0 {
		return errOutOfStock
	}
	cd.Stock--
	return nil
}
