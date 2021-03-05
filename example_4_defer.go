// go has a garbage collector which means you don't have to deal with memory management when allocate object and stop using it
// Go's garbage collector will clear it up
// However, memory is just one kind of resource. And you may have other resources
// you use in your program. For example, files, sockets, virtual machines and others.
// You'd like to make sure that these resources are closed when you're done with them as well.
// To make sure a resource is closed, use defer.

package main

import (
	"fmt"
)

func cleanup(name string) {
	fmt.Printf("Cleaning up %s \n", name)
}

// defer the code in reverse order
func worker() {
	defer cleanup("A")
	defer cleanup("B")
	fmt.Println("worker")
}

func main() {
	worker()
}
