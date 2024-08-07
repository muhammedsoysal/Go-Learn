package string_functions

//alias
import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "Muhammed"
	fmt.Println(s.Count(isim, ""))
	fmt.Println(s.Count(isim, "m"))
	fmt.Println(s.Contains(isim, "c"))
	fmt.Println(s.Index(isim, "m"))
	
}
