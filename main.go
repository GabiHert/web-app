package main

import (
	"net/http"
	"web-app/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.RenderRoutes()
	http.ListenAndServe(":8000", nil)

}
