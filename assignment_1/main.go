package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Greeting() {
	fmt.Printf("Hello, %s!", p.Name)
}

type Employee struct {
	Name string
	ID   int
}

type Manager struct {
	Employee
	Department string
}

func (e *Employee) Work() {
	fmt.Printf("Working on %s, %d", e.Name, e.ID)
}


type Circle struct {
	r float64
}

type Rectangle struct {
	a float64
	b float64
}

type Shape interface {
	area() float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	return r.a * r.b
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.area())
}


type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func encode(p Product) string {
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Error encoding to JSON: %v", err)
	}
	return string(jsonData)
}

func decode(jsonStr string) Product {
	var p Product
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	return p
}

func main() {
	fmt.Println("Hello World!")

	var a int32
	a = 12
	fmt.Printf("Integer a = %d\n", a)

	var b float64
	b = 1.2
	fmt.Printf("Float b = %f\n", b)

	c := false
	fmt.Printf("Boolean = %t\n", c)

	sum := 0
	for i := 0; i < 10; i++ {
		if isPositive(i) {
			sum += i
		}
	}
	fmt.Printf("Sum = %d\n", sum)

	const Weekday int = 2
	switch Weekday {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Thursday")
	}

	fmt.Printf("Sum = %d\n", add(2, 3))
	swap(2, 3)
	fmt.Println(quotientRemainder(8, 5))

	person := Person{Name: "Azamat", Age: 21}
	person.Greeting()

	mgr := Manager{
		Employee: Employee{
			Name: "assdasd", ID: 10},
		Department: "IT",
	}

	mgr.Work()

	circle := Circle{r: 3}
	rectangle := Rectangle{a: 3, b: 4}

	PrintArea(&circle)
	PrintArea(&rectangle)

	product := Product{Name: "Iphone", Price: 999.99, Quantity: 10}

	jsonStr := encode(product)
	fmt.Println("Product in JSON:", jsonStr)

	decodedProduct := decode(jsonStr)
	fmt.Printf("Decoded Product: %+v\n", decodedProduct)
}

func isPositive(a int) bool {
	if a > 0 {
		return true
	}
	return false
}


func add(a, b int) (sum int) {
	sum = a + b
	return sum
}

func swap(a, b int) {
	fmt.Printf("Before swap a = %d, b = %d\n", a, b)
	a = b
	b = a
	fmt.Printf("After swap a = %d, b = %d\n", a, b)
}

func quotientRemainder(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return quotient, remainder
}