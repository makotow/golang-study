package examples

//https://gobyexample.com/pointers
//
func zeroval(ival int) {
  ival = 0
}

func zeroptr(iptr *int) {
  // * means dereferences the pointer
  *iptr = 0
}
//
//func main() {
//  i := 1
//  fmt.Println("Initial:", i)
//
//  zeroval(i)
//  fmt.Println("zeroval:", i)
//
//  // &i syntax gives the memory address of i
//  zeroptr(&i)
//  fmt.Println("zeroptr:", i)
//
//  fmt.Println("pointer:", &i)
//}
