package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"

	"cli/db"
)

const CATEGORY_PREFIX = "Category"
const PRODUCT_PREFIX = "Product"

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate test data",
	Run: func(cmd *cobra.Command, args []string) {
		defer db.GetConn().Close()
		var waitgroup sync.WaitGroup
		for i := 1; i <= 200; i++ {
			waitgroup.Add(1)
			go insertCategory(
				fmt.Sprintf("%v %v", CATEGORY_PREFIX, i),
				&waitgroup,
			)
		}
		waitgroup.Wait()

		categoryId := 1
		count := 0
		for i := 1; i <= 100000; i++ {
			waitgroup.Add(1)
			go insert(
				fmt.Sprintf("%v %v", PRODUCT_PREFIX, i),
				categoryId,
				&waitgroup,
			)
			count++
			if count == 500 {
				count = 0
				categoryId++
			}
		}
		waitgroup.Wait()
	},
}

func insert(productName string, categoryId int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	query := fmt.Sprintf(
		"INSERT INTO product(fk_category, name, image) VALUES ('%v', '%v', '%v')",
		categoryId,
		productName,
		"https://via.placeholder.com/350x150",
	)
	db.GetConn().MustExec(query)
}

func insertCategory(categoryName string, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	query := fmt.Sprintf(
		"INSERT INTO category(name, image) VALUES ('%v', '%v')",
		categoryName,
		"https://via.placeholder.com/350x150",
	)
	db.GetConn().MustExec(query)
}
