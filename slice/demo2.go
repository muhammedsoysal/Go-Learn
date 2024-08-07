package slices

import "fmt"

func Demo2() {
	citys := []string{"Ankara", "Istanbul", "Izmır"}
	fmt.Println(citys)

	citysCopy := make([]string, len(citys))
	fmt.Println(citysCopy)
	copy(citysCopy, citys)
	fmt.Println(citysCopy)

	fmt.Println(citys[0:2])
	fmt.Println(citys[1:])
	fmt.Println(citys[:3])

}
