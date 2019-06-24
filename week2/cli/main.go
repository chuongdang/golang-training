package main

import (
	"cli/cmd"
	"cli/db"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)
	cmd.Execute()
	db.GetConn().Close()
}
