package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"M3", 50000, "Apple"})
	fmt.Println(product{name: "M3", unitPrice: 50000})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
