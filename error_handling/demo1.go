package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo5.txt")

	//nil
	if err != nil {
		//type assertion
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı: ", pErr.Path)
			return
		}
		fmt.Println("Dosya Bulunamadı")

	} else {
		fmt.Println(f.Name())
	}
}
