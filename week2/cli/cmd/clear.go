package cmd

import (
	"github.com/spf13/cobra"

	"cli/db"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear test data",
	Run: func(cmd *cobra.Command, args []string) {
		query := "TRUNCATE TABLE product"
		db.GetConn().MustExec(query)
	},
}
