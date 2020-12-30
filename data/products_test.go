package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "Tea",
		Price: 2.0,
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
