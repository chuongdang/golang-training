package main

import (
	"fmt"
	"product-service/routes"
)

func main() {
	fmt.Println(db.DB)
	routes.Start()
}
