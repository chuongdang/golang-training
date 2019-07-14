package main

import (
	"flag"
	"fmt"
)

func main() {
	var source string
	var dest string

	flag.StringVar(&source, "source", "", "Source of moving")
	flag.StringVar(&dest, "dest", "", "Destination to move to")
	flag.Parse()
	moveSomething(&source, &dest)
}

func moveSomething(source *string, dest *string) {
	fmt.Println("Moving from " + *source + " to " + *dest + "...")
}