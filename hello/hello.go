package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Printf("hello, world\n")
	/*
			vertices := make(map[string]int)
			vertices["triangle"] = 2
			vertices["square"] = 3
			vertices["dodecagon"] = 12

			fmt.Println(vertices["triangle"])

			for i := 0; i < 5; i++ {
				if i < 2 {
					fmt.Println("<2")
				} else {
					fmt.Println(i)
				}
			}

		result := sum(2, 3)
		fmt.Println(result)

		sqrt, err := sqrt(4)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(sqrt)
		}
	

	p := person{name: "Jake", age: 23}
	fmt.Println(p)
	fmt.Println("Person name: ", p.name)
		*/
	i := 8
	inc(*i)
	fmt.Println(i)
}

func inc(x *int) {
	*x++
}

type person struct {
	name string
	age  int
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}
