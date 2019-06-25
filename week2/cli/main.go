package main

import (
	"cli/cmd"
	"cli/db"
)

func main() {
	cmd.Execute()
	db.GetConn().Close()
}
