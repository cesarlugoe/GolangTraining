package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + p.last
}

//(p person) refers to the reciever, so 'person' is the type struct or 'Class' and 'p' is refers to the
// instance of thar struct or class in th ex. below (p1 and p2)

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName()) //sugarized version of Person.fullName(p2)

}
