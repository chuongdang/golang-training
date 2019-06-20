package main

import "webservice/routers"
import "webservice/db"

func main() {
	defer db.GetConn().Close()
	routers.Run("9191")
}
