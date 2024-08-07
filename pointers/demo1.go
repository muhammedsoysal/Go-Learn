package pointers

import "fmt"

func Demo1(sayi *int) {
	*sayi++
	fmt.Println("Demodaki sayı: ", sayi)
}
