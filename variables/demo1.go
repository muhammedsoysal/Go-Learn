package variables

import "fmt"

func Demo1() {
	var helloWorld string = "helloWorld"

	fmt.Println(helloWorld)
	fmt.Println(helloWorld)
	fmt.Println(helloWorld)
	fmt.Println(helloWorld)

	kdv := 25
	s := "test"
	f := 2.5

	fmt.Println(kdv)
	fmt.Printf("veri tipi: %T\n", kdv)
	fmt.Println(s)
	fmt.Printf("veri tipi: %T\n", s)
	fmt.Println(f)
	fmt.Printf("veri tipi: %T\n", f)

	var durum bool

	var metin1 string = "Engin"
	var metin2 string = "Soysal"

	durum = metin1 == metin2

	fmt.Println(durum)
	fmt.Println(!durum)
}
