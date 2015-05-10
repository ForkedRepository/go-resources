// maps
package main

import (
	"fmt"
)

func main() {
	// make() and new()
	// http://golang.org/doc/effective_go.html#allocation_new
	// http://stackoverflow.com/a/9325620/2536357
	x := make(map[string]int)
	x["key"] = 10
	if name, ok := x["key"]; ok { // check if map[key] exists
		// map[key] returns name, ok
		// if not in map: name = 0, ok = false
		// else: name = value, ok = true
		fmt.Println(name, ok)
	}
	fmt.Println(x)
}
