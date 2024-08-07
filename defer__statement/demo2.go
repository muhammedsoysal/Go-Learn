package defer_statement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()

	if sayi%2 == 0 {
		return "Çift Sayıdır"
	}
	if sayi%2 != 0 {
		return "Tek Sayıdır"
	}

	return "Belli Değil"
}
func Test() {
	fmt.Println("Sonuc: ", Demo2(10))
}

func DeferFunc() {
	fmt.Println("Biiti")
}
