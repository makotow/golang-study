package main

// define interface
type SomeBehivor interface {
	DoSomething(arg string) string
}

// define struct
type StructA struct {
}

func (self *StructA) DoSomething(arg string) string {
	return arg
}

func main() {
	var behivor SomeBehivor
	behivor = &StructA{}
	behivor.DoSomething("A")

}
