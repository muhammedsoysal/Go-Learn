package for_range

import "fmt"

//1-10 arasındaki sayılardan tek olanları toplayan eleman
//for-range
func Demo2() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
 
	total := 0
	for _, number := range numbers {
		if number%2 != 0 {
			total += number
		}
	}
	fmt.Println(total)
}
