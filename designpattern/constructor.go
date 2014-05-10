// 写経　http://blog.monochromegane.com/blog/2014/03/23/struct-implementaion-patterns-in-golang/
package main

import "fmt"

type StructA struct {
	Name string
}

func (self *StructA) SomeInitialize() {
	// initialize
}

// define constructor function
func NewStructA(name string) *StructA {
	structA := &StructA{Name: name}
	structA.SomeInitialize()
	return structA
}

func main() {
	s := NewStructA("name")
	fmt.Println(s.Name)
}
