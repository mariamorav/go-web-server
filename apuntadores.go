package main

import "fmt"

func main() {
	x := 25

	fmt.Println(x)
	fmt.Println(&x)
	changeValue(x)
	y := &x
	fmt.Println(y)
	fmt.Println(*y)

}

func changeValue(a int) {
	fmt.Println(&a)
	a = 33
}
