package controller

import (
	authcontroller "recsvc/internal/controller/auth"
	maincontroller "recsvc/internal/controller/main"
	_ "recsvc/internal/controller/middleware" // avoid import cycle
)

type Handlers struct {
	Auth *authcontroller.Handler
	Main *maincontroller.Handler
}

func NewHandlers(
	auth *authcontroller.Handler,
	main *maincontroller.Handler,
) *Handlers {
	return &Handlers{
		Auth: auth,
		Main: main,
	}
}
