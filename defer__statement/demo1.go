package defer_statement

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalışt.")
}
func C() {
	fmt.Println("C fonksiyonu çalışt.")
}
func D() {
	fmt.Println("D fonksiyonu çalışt.")
}
func B() {
	defer A()
	defer D()
	defer C()
	fmt.Println("B fonksiyonu çalışt.")
}

