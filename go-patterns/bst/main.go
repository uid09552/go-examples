package main

import (
	"fmt"
	bst "uid09552/go-examples/go-patterns/bst/node"
)

func main() {
	root := bst.NewTree()
	fmt.Printf("%v", root)
	root.Add(10)
	root.Add(20)
	root.Add(15)
	root.Add(5)
	root.Add(6)
	root.Add(50)
	root.Add(55)
	root.Add(58)
	root.Add(57)
	root.Add(59)
	root.Add(60)
}
