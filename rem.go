package main

import "fmt"

func main() {
	p := newPerson(
		name("panjinhao"),
		age(19),
		sex("male"),
	)
	fmt.Println(p)
}

type Person struct {
	name string
	age  int
	sex  string
}

func newPerson(options ...option) Person {
	p := &Person{}
	for _, option := range options {
		option(p)
	}
	return *p
}

type option func(*Person)

func name(name string) option {
	return func(p *Person) {
		p.name = name
	}
}

func age(age int) option {
	return func(p *Person) {
		p.age = age
	}
}
func sex(sex string) option {
	return func(p *Person) {
		p.sex = sex
	}
}
