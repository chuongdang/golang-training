package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"cli/db"
)

var generateNormalCmd = &cobra.Command{
	Use:   "generate-normal",
	Short: "Generate test data (without goroutine)",
	Run: func(cmd *cobra.Command, args []string) {
		defer db.GetConn().Close()
		defaultProductName := "Product"
		for i := 0; i < 1000000; i++ {
			query := fmt.Sprintf(
				"INSERT INTO product(name) VALUES ('%v %v')",
				defaultProductName,
				i,
			)
			db.GetConn().MustExec(query)
		}
	},
}
