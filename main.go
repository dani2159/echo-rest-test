package main

import (
	"github.com/dani2159/echo-rest/databases"
	"github.com/dani2159/echo-rest/routes"
)

func main() {
	databases.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
