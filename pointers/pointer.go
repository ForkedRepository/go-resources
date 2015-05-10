// pointer
package main

import (
	"fmt"
)

type User struct {
	Name string
}

func main() {
	u := &User{Name: "Leto"}
	fmt.Println("this value is copied into the function (standard behaviour):")
	fmt.Println(u.Name)
	Modify(u)
	fmt.Println(u.Name)

	x := 5
	fmt.Println("this value is dereferenced and modified (& and *):")
	fmt.Println(x)
	zero(&x)
	fmt.Println(x)
}

func zero(xPtr *int) {
	*xPtr = 0
}

func Modify(u *User) {
	u.Name = "Paul"
}

/*
http://www.golang-book.com/8/index.htm
The * and & operators

In Go a pointer is represented using the * (asterisk) character followed by the
type of the stored value. In the zero function xPtr is a pointer to an int.

* is also used to “dereference” pointer variables. Dereferencing a pointer gives
us access to the value the pointer points to. When we write *xPtr = 0 we are
saying “store the int 0 in the memory location xPtr refers to”.
If we try xPtr = 0 instead we will get a compiler error because xPtr is not an
int it's a *int, which can only be given another *int.

Finally we use the & operator to find the address of a variable. &x returns
a *int (pointer to an int) because x is an int. This is what allows us to
modify the original variable. &x in main and xPtr in zero refer to the same
memory location.
*/
