package services

import (
	"encoding/json"
	"io"
	"net/http"
	"webservice/controllers"
)

func RegisterControllers() {
	uc := controllers.NewUserController()

	// matching '/users' pattern, handle with uc created
	http.Handle("/users", uc)

	// '/users/*' should be handled there as well
	http.Handle("/users/", uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
