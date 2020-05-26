package main

import "fmt"
import "../utils"

var lst = []int{1, 9, 3, 2, 6, 5, 8, 7, 4}

func main() {
	for i := 0; i < len(lst)-1; i++ {
		for j := 1; j < len(lst); j++ {
			if lst[j] < lst[j-1] {
				lst[j], lst[j-1] = utils.Swap(lst[j], lst[j-1])
			}
		}
	}

	fmt.Println(lst)

}
