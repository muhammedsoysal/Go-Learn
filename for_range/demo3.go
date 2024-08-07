package for_range

import "fmt"

func Demo3() {
	sozluk := map[string]string{"Glass": "Bardak", "microphone": "mikroffon"}

	for k := range sozluk {
		fmt.Println(k)
	}
}
