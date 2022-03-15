package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWarehouse(t *testing.T) {
	cd := CD{Name: "Physical Graffiti", Artist: "Led Zeppelin", Stock: 1}

	t.Run("initialize a warehouse", func(t *testing.T) {
		w := NewWarehouse()
		assert.Len(t, w.CDs, 0)
	})
	t.Run("add a CD", func(t *testing.T) {
		w := Warehouse{}
		w.AddCD(&cd)
		assert.Len(t, w.CDs, 1)
	})
	t.Run("search CD by name", func(t *testing.T) {
		w := Warehouse{}
		w.AddCD(&cd)
		assert.Len(t, w.CDs, 1)
		lookupCD := CD{Name: "Physical Graffiti", Artist: ""}
		results := w.FindCD(&lookupCD)
		assert.Len(t, results, 1)
	})
	t.Run("search CD by artist", func(t *testing.T) {
		w := Warehouse{}
		w.AddCD(&cd)
		assert.Len(t, w.CDs, 1)
		lookupCD := CD{Name: "", Artist: "Led Zeppelin"}
		results := w.FindCD(&lookupCD)
		assert.Len(t, results, 1)
	})
	t.Run("search CD by artist and name", func(t *testing.T) {
		w := Warehouse{}
		w.AddCD(&cd)
		assert.Len(t, w.CDs, 1)
		lookupCD := CD{Name: "Physical Graffiti", Artist: "Led Zeppelin"}
		results := w.FindCD(&lookupCD)
		assert.Len(t, results, 1)
	})
	t.Run("search CD, nonexistent result", func(t *testing.T) {
		w := Warehouse{}
		w.AddCD(&cd)
		assert.Len(t, w.CDs, 1)
		lookupCD := CD{Name: "", Artist: "The Clash"}
		results := w.FindCD(&lookupCD)
		assert.Len(t, results, 0)
	})
}
func TestCustomer(t *testing.T) {
	t.Run("buy CD", func(t *testing.T) {
		cd := CD{Name: "Physical Graffiti", Artist: "Led Zeppelin", Stock: 1}
		w := Warehouse{}
		w.AddCD(&cd)
		assert.Len(t, w.CDs, 1)
		lookupCD := CD{Name: "Physical Graffiti", Artist: "Led Zeppelin"}
		results := w.FindCD(&lookupCD)
		assert.Len(t, results, 1)
		var result = results[0]
		// Verify that stock is available:
		assert.Equal(t, result.Stock, 1)

		// Buy the CD:
		customer := Customer{}
		err := customer.Buy(result)
		assert.NoError(t, err)

		// Verify that stock is decreased:
		assert.Equal(t, result.Stock, 0)
	})
	t.Run("buy CD (out of stock)", func(t *testing.T) {
		cd := CD{Name: "Physical Graffiti", Artist: "Led Zeppelin", Stock: 0}
		w := Warehouse{}
		w.AddCD(&cd)
		assert.Len(t, w.CDs, 1)
		lookupCD := CD{Name: "Physical Graffiti", Artist: "Led Zeppelin"}
		results := w.FindCD(&lookupCD)
		assert.Len(t, results, 1)
		var result = results[0]
		customer := Customer{}

		// Tries to buy a CD that's out of stock:
		err := customer.Buy(result)
		assert.EqualError(t, err, "requested CD is out of stock")
	})
	t.Run("leave CD review", func(t *testing.T) {})
}

func CustomerLeaveReview(t *testing.T) {
	// customer := &Customer{}
	// customer.LeaveReview(*cd*)
}

func RecordLabelAddCDs(t *testing.T) {}
