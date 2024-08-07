package errorhandling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 50

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz")
	}

	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayi giriniz", nil
	} else if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayi giriniz", nil
	}

	return "Bildiniz", nil
}

func Demo2() {

	mesaj, err := TahminEt(50)

	fmt.Println(mesaj)
	fmt.Println(err)

}
