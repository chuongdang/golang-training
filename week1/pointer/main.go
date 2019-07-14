package main

import "fmt"

func main() {
	name := "Chuong"
	helloName(&name)
	fmt.Println(name)
}

func helloName(name *string) {
	*name = "Hello " + *name
}
