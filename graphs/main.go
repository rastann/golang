package main

import "fmt"

func main() {
	fmt.Println("Hello razvan")

	m := make(map[int]int)
	m[1] = 12
	m[2] = 14

	display(m)

	mm := map[int]int{1: 434, 2:3243}
	display(mm)
}

func display(m map[int]int) {
	for k,v := range m {
		fmt.Printf("key=%d, value=%d", k, v)
	}
}
