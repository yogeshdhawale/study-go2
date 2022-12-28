package main

import "fmt"

func mapeg() {
	grades := make(map[string]float32)
	grades["A"] = 80
	grades["B"] = 70
	grades["C"] = 60

	fmt.Println("Map:")
	fmt.Println(grades)

	for x, y := range grades {
		fmt.Printf("%s got %g\n", x, y)
	}
	fmt.Println()
	
	delete(grades, "A")
	fmt.Println(grades)
}
