package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	// Starting to think what this could be, it gets the address of the Struct
	// Allocates a new space in memory of the value at that address
	p := &Product{
		Name:  "Danno",
		Price: 1.00,
		SKU:   "abs-abc-def",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
