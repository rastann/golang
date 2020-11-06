package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)

	l, v := 0, 0
	for _, val := range s {
		if val == 'U' {
			l++
		}
		if val == 'D' {
			l--
		}

		if l == 0 && val == 'U' {
			v++
		}
	}

	fmt.Println("Nr of valleys: ", v)
}