package main

import (
	"generator/db"
	"fmt"
	"sync"
)

func main() {
	defer db.GetConn().Close()
	defaultProductName := "Product"
	var waitgroup sync.WaitGroup
	waitgroup.Add(100)
	for i := 0; i < 100; i++ {
		go insert(
			fmt.Sprintf("%v %v", defaultProductName, i),
			&waitgroup,
		)
	}
	waitgroup.Wait()
}

func insert(productName string, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	query := fmt.Sprintf(
		"INSERT INTO product(name) VALUES ('%v')",
		productName,
	)
	db.GetConn().MustExec(query)
}
