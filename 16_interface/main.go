package main

import "fmt"

// ============================================
// EXAMPLE 1: Basic Interface with Multiple Types
// ============================================

// Speaker interface - any type that has a speak() method automatically implements this
type speaker interface {
	speak() string
}

// Cat type
type cat struct {
	name string
}

// Cat implements speaker interface (no explicit declaration needed!)
func (c cat) speak() string {
	return "Meow! My name is " + c.name
}

// Dog type
type dog struct {
	name  string
	breed string
}

// Dog also implements speaker interface
func (d dog) speak() string {
	return "Woof! I'm " + d.name + ", a " + d.breed
}

// Human type
type human struct {
	name string
}

// Human also implements speaker interface
func (h human) speak() string {
	return "Hello, I'm " + h.name
}

// This function accepts ANY type that implements speaker interface
// This is POLYMORPHISM - different types, same interface
func makeSound(s speaker) {
	fmt.Println(s.speak())
}

// ============================================
// EXAMPLE 2: Interface with Multiple Methods
// ============================================

// Animal interface requires both speak() and move() methods
type animal interface {
	speak() string
	move() string
}

// Bird type
type bird struct {
	species string
}

// Bird implements animal interface (has both methods)
func (b bird) speak() string {
	return "Chirp chirp!"
}

func (b bird) move() string {
	return "Flying in the sky"
}

// Fish type
type fish struct {
	species string
}

// Fish implements animal interface
func (f fish) speak() string {
	return "Blub blub..."
}

func (f fish) move() string {
	return "Swimming in water"
}

// Function that works with any animal
func describeAnimal(a animal) {
	fmt.Println("Sound:", a.speak())
	fmt.Println("Movement:", a.move())
}

// ============================================
// EXAMPLE 3: Real-World Use Case - Shapes
// ============================================

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14159 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14159 * c.radius
}

// Calculate total area of any shapes
func printShapeInfo(s shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.area(), s.perimeter())
}

func main() {
	fmt.Println("=== EXAMPLE 1: Basic Interface ===")
	// Different types, all implement speaker
	myCat := cat{name: "Whiskers"}
	myDog := dog{name: "Buddy", breed: "Golden Retriever"}
	me := human{name: "Ankush"}

	// Same function works with all types!
	makeSound(myCat)
	makeSound(myDog)
	makeSound(me)

	fmt.Println("\n=== EXAMPLE 2: Multiple Methods ===")
	// Types implementing animal interface
	sparrow := bird{species: "Sparrow"}
	goldfish := fish{species: "Goldfish"}

	describeAnimal(sparrow)
	fmt.Println()
	describeAnimal(goldfish)

	fmt.Println("\n=== EXAMPLE 3: Shapes (Real-World) ===")
	rect := rectangle{width: 10, height: 5}
	circ := circle{radius: 7}

	fmt.Print("Rectangle - ")
	printShapeInfo(rect)

	fmt.Print("Circle - ")
	printShapeInfo(circ)

	// Store different shapes in a slice
	fmt.Println("\nCalculating total area of all shapes:")
	shapes := []shape{rect, circ}
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.area()
	}
	fmt.Printf("Total area: %.2f\n", totalArea)

	fmt.Println("\n=== KEY CONCEPTS ===")
	fmt.Println("1. Interfaces are IMPLICIT - no 'implements' keyword needed")
	fmt.Println("2. If a type has all methods of an interface, it implements it")
	fmt.Println("3. One function can work with many different types (polymorphism)")
	fmt.Println("4. Interfaces enable flexible, testable code")
}
