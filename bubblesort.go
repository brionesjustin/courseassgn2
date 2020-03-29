package main

import (
	"fmt"
)

var inputinte int


func bubblesort(sli []int) {
	n := len(sli)
	swapped := true
	for swapped {
		swapped = false 
		for i := 1; i < n; i++ {
			if sli[i] < sli[i-1] {
				sli[i], sli[i-1] = sli[i-1], sli[i]
				swapped = true
			}
		}
}
}

//func main() {
	a := make([]int, 0)
	for {
		fmt.Println("Please enter an integer:")
		_, err := fmt.Scan(&inputinte)
		if err != nil {
			fmt.Println(err)
		}
		if len(a) <= 10 {
			a = append(a, inputinte)
		} else {
			break
		}	
	}
	//fmt.Println(a)
	bubblesort(a)
	fmt.Println(a)
}