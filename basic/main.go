package main

import (
	"errors"
	"fmt"

	hello "githuv.com/apparatno/go-workshop/module"
)

type Thing string

type person struct {
	name string
	age  int
}

var array = []int{1, 2, 3, 4, 5, 6}

func main() {
	gree := hello.Hello()

	fmt.Println(gree)
	fmt.Printf("Hello %s\n", "World")
	a := 1
	b := 2
	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d\n", sumList(array))

	p := person{
		"Kristian",
		28,
	}

	// Pointers

	point := &p

	thing(point)

	greeting, err := formatPerson(p)

	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)
	fmt.Println(p.format())
	slice := make([]person, 0, 10)
	slice = append(slice, person{name: "Me", age: 3})

	for _, p2 := range slice {
		formatP2, _ := formatPerson(p2)
		fmt.Println(formatP2)
	}

	aMap := make(map[string]person)

	aMap["foo"] = person{name: "You", age: 2}

	_, ok := aMap["bar"]
	if !ok {
		panic("YAAAAA")
	}

	fmt.Printf("length: %d\n", len(slice))


	// Interfaces!

	p.work()
}



func (p person) work() {
	p
}

func formatPerson(p person) (string, error) {
	if p.name == "" {
		return "", errors.New("no name!")
	}
	return fmt.Sprintf("Hello %s!", p.name), nil
}

func (p person) format() string {
	return fmt.Sprintf("Hello %s!", p.name)
}

func thing(p *person) {
	p.name = "DUMB"
}
