package main

import (
	"fmt"
)

var(
	list1, list2 []int
	n1, n2	       int
)

func main() {
	fmt.Println("Start entering the things fot the same plqwwzzzz: ")
	for {
		_, err := fmt.Scan(&n1, &n2)
		if err != nil {
			break
		}
		append(list1, &n1)
		append(list2, &n2)
	}
	fmt.Println(list1, list2)
}
