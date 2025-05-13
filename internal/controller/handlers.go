package controller

import (
	authcontroller "glossika/internal/controller/auth"
	maincontroller "glossika/internal/controller/main"
	_ "glossika/internal/controller/middleware" // avoid import cycle
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
