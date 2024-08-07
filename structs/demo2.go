package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) Save() {
	fmt.Println("Eklendi:", c)
}

func (c customer) Update() {
	fmt.Println("Güncellendi:", c)
}
func (c customer) Delete() {
	fmt.Println("Silindi:", c)
}

func Demo2() {
	c := customer{"Muhammed", "Soysal", 19}
	c.Save()
	c.Update()
	c.Delete()
}
