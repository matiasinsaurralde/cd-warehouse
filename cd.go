package main

/*

Customer side

1) Search for CDs by artist name or album name.
2) Buy CDs using a credit card through our external payment provider.
3) Leave reviews (rating + comment) for each CD that was bought.

Warehouse side

1) Add CDs and stock availability for material that's sent to the CD warehouse.

*/

type CD struct {
	Name   string
	Artist string
	Stock  int
}
