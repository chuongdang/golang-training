package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"

	"cli/db"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate test data",
	Run: func(cmd *cobra.Command, args []string) {
		defer db.GetConn().Close()
		defaultProductName := "Product"
		var waitgroup sync.WaitGroup
		waitgroup.Add(1000000)
		for i := 0; i < 1000000; i++ {
			go insert(
				fmt.Sprintf("%v %v", defaultProductName, i),
				&waitgroup,
			)
		}
		waitgroup.Wait()
	},
}

func insert(productName string, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	query := fmt.Sprintf(
		"INSERT INTO product(name) VALUES ('%v')",
		productName,
	)
	db.GetConn().MustExec(query)
}
