package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
  fmt.Println(s.age)

  ss := s
  fmt.Println(ss.age)
  ss.age = 11

  fmt.Println(ss.age)
  fmt.Println(s.age)
}
