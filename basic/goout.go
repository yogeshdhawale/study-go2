package main

import (
	"fmt"
	"math"
	"math/rand"
)

func testmath() {

	fmt.Println("Absolute value:", math.Floor(3.4))
	fmt.Println("sqrt of 4 is:", math.Sqrt(4))

	rand.Seed(0)
	fmt.Println("Random number:", rand.Intn(100))
}

func add(x float64, y float64) float64 {
	return x + y
}

const PI float64 = 3.14

func concat1(a, b string) (string, string) {
	return a, b
}

func gotypes() {

	//var x float64 = 2.1
	//var y float64 = 3.1

	x, y := 2.2, 3.2

	fmt.Printf("Adding: %g + %g = %g\n", x, y, add(x, y))

	w1, w2 := "Hi", "there"

	fmt.Print("concating:")
	fmt.Println(concat1(w1, w2))

	var b float64 = float64(5)
	fmt.Println(b)
}

func gorefs() {
	x := 15
	y := &x
	fmt.Println("value is:", x)
	fmt.Println("Ref is:", y)
	fmt.Println("Ref * is:", *y)
	*y = 5
	fmt.Println("value after changing *y is:", x)
	//squre
	*y = *y * *y
	fmt.Println("Square is:", *y)
}

// struct
type car struct {
	gas_p     uint16
	break_p   uint16
	wheel     int16
	top_speed float64
}

const PEDALMAX float64 = 65535
const km2mi = 1.60934

func (c car) kmh() float64 {
	return float64(c.gas_p) * (c.top_speed / PEDALMAX)
}

func (c car) mph() float64 {
	return float64(c.gas_p) * (c.top_speed / PEDALMAX) / km2mi
}

func (c *car) updateSpeed(x float64) {
	c.top_speed = x
}
func main() {
	fmt.Println("Hello world 1 ...")
	fmt.Println()

	//testmath()

	//gotypes()

	//gorefs()

	v1 := car{65000, 0, 12341, 225}
	//v2 := car{gas_p: 1, break_p: 2, wheel: 3, top_speed: 4}

	fmt.Println("Gas pedal:", v1.gas_p)
	fmt.Println("km per hr:", v1.kmh())
	fmt.Println("mi per hr:", v1.mph())

	v1.updateSpeed(500)
	fmt.Println("new speed:", v1.top_speed)

	fmt.Println("\n***Goodbye ...")
}
