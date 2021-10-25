package main

import (
	"kudos/endpoints/routes"
)

func main() {
	// Load routes
	e := routes.LoadRoutes()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
