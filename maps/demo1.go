package maps

import "fmt"

func Demo1() {

	//key-value
	sozluk := make(map[string]string)
	sozluk["Table"] = "Masa"
	sozluk["Book"] = "Kitap"
	sozluk["Pencil"] = "Kalem"

	fmt.Println(sozluk["Pencil"])
	fmt.Println("Sözlük Eleman Sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)
	delete(sozluk, "Book")

	fmt.Println("Sözlük Eleman Sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)

	deger, varMi := sozluk["Pencil"]
	fmt.Println(deger)
	fmt.Println("Listede Olma Durumu:", varMi)

	sozluk2 := map[string]string{"Glass": "Bardak", "microphone": "mikroffon"}

	fmt.Println(sozluk2)

	


}
