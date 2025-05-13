package api

import (
	controller "glossika/internal/controller"
	"glossika/internal/controller/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAPIHandlers(r *gin.Engine, h *controller.Handlers) {
	r.Use(middleware.Cors())
	r.Use(middleware.ErrorHttp)

	r.POST("/login", h.Auth.Login)
	r.POST("/sign-up", h.Auth.SignUp)
	r.POST("/confirm-sign-up", h.Auth.ConfirmSignUp)

	pR := r.Group("/")
	pR.Use(middleware.Jwt())
	{
		pR.GET("/recommendation", h.Main.GetRecommendation)
	}
}
