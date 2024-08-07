package main

import (
	"fmt"
	"goLearn/project"
)

func main() {

	// variables.Demo1()
	// fmt.Println()
	// slices.Demo1()
	// slices.Demo2()
	// arr := []int{1, 32, 45, 8}
	// fmt.Println(functions.ToplaVariadic(1, 2, 4, 5, 9, 10))
	// fmt.Println(functions.ToplaVariadic(1, 2))
	// fmt.Println(functions.ToplaVariadic(arr...))
	// fmt.Println(functions.ToplaVariadic())

	// maps.Demo1()

	// for_range.Demo1()
	// for_range.Demo2()
	// for_range.Demo3()

	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("Maindeki sayi: ", sayi)

	// sayilar := []int{1, 2, 3}
	// pointers.Demo2(sayilar)
	// fmt.Println("Maindaki sayilar: ", sayilar)

	// structs.Demo2()

	// go goroutines.TekSayilar()
	// go goroutines.CiftSayilar()
	// time.Sleep(5 * time.Second)
	// fmt.Println("Main Bitti")

	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)
	// go channels.TekSayilar(tekSayiCn)
	// go channels.CiftSayilar(ciftSayiCn)

	// ciftSayiToplam, TektekSayiToplam := <-ciftSayiCn, tekSayiCn

	// carpim := ciftSayiToplam * <-TektekSayiToplam
	// fmt.Println("Çarpım: ", carpim)

	// 	interfaces.Demo1()
	// 	interfaces.Demo2()

	// defer_statement.B()
	// defer_statement.Test()
	// defer_statement.Demo3()
	// defer_statement.A()

	// errorhandling.Demo1()
	// errorhandling.Demo2()
	// errorhandling.Demo3()

	// string_functions.Demo1()

	// restfull.Demo1()
	// restfull.Demo2()

	// interfaces.Demo3()

	products, _ := project.GetAllProducts()

	for _, p := range products {
		fmt.Println(p.ProductName)
	}

	product, _ := project.AddProduct()

	fmt.Println(product)
}
