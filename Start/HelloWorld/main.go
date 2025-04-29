// package main

// import "fmt"
// import "os"
// func main() {
// 	fmt.Println("Hello world")
// 	os.Exit(0)
// }

package main

import "fmt"

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("HelloWorld"[1])
	fmt.Println("Hello " + "World")

	fmt.Println(31132 * 42452)
	fmt.Println((true && false) || (false && true) || !(false && false))
}
