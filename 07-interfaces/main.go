package main

import (
	"fmt"
	"math"
)

// In Go, interfaces are satisfied IMPLICITLY!
// A type does NOT declare "implements Shape".
// If a type implements Area() and Perimeter(), it automatically IS a Shape!
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle type
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle type
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Polymorphic function accepting ANY type that implements Shape
func PrintShapeInfo(s Shape) {
	fmt.Printf("Shape -> Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// The `any` type (alias for `interface{}`) can hold any value whatsoever
func DescribeAnything(val any) {
	// Type switch: checks the dynamic type of an interface
	switch v := val.(type) {
	case int:
		fmt.Printf("Integer: %d (multiplied: %d)\n", v, v*2)
	case string:
		fmt.Printf("String: %q (length: %d)\n", v, len(v))
	case Shape:
		fmt.Printf("A Shape with Area: %.2f\n", v.Area())
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	fmt.Println("=== 07: Interfaces and Polymorphism ===")

	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	// Polymorphism in action
	PrintShapeInfo(rect)
	PrintShapeInfo(circ)

	// Slice of interfaces
	shapes := []Shape{rect, circ}
	for i, s := range shapes {
		fmt.Printf("Shape %d Area: %.2f\n", i+1, s.Area())
	}

	// Any and Type switches
	fmt.Println("\nTesting type switches:")
	DescribeAnything(42)
	DescribeAnything("Hello, Go!")
	DescribeAnything(rect)
}
