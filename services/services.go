package services

import (
	"fmt"
	"github.com/arbor-dev/arbor"
	"net/http"
)

var Routes = arbor.RouteCollection{
	arbor.Route{
		"Gateway",
		"GET",
		"/",
		Gateway,
	},
}

func Gateway(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The API Gateway Lives")
}

func RegisterAPIs() arbor.RouteCollection {
	Routes = append(Routes, TestRoutes...)
	Routes = append(Routes, AuthRoutes...)
	Routes = append(Routes, UserRoutes...)
	Routes = append(Routes, RegistrationRoutes...)
	return Routes
}
