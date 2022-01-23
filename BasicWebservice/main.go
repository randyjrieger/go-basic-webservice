package main

import (
	"net/http"
	"webservice/services"
)

func main() {
	services.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
