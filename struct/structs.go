// structs
package main

import (
	"fmt"
)

// An interface is a structure for methods common to different structures
// we define a “method set”. A method set is a list of methods that a type
// must have in order to “implement” the interface.
// http://www.golang-book.com/9/index.htm#section2
type Being interface {
	Talk()
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// To describe Has-A relatioship
type Person struct {
	Name string
}

// To describe Is-A relationship
type Android struct {
	Person
	Model string
}

func main() {
	x := Person{"Joy"}
	x.Talk()

	a := new(Android)
	a.Person.Name = "Rick"
	a.Model = "Jc4"
	a.Person.Talk()
}
