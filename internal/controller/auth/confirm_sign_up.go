package authcontroller

import (
	"net/http"

	"recsvc/internal/domain/er"
	authucase "recsvc/internal/usecase/auth"
	"recsvc/internal/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ConfirmSignUp(c *gin.Context) {
	var req authucase.ConfirmSignUpIn
	if !util.ParseValidate(c, &req) {
		return
	}

	if err := h.Usecase.ConfirmSignUp(c.Request.Context(), req); err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.Status(http.StatusOK)
}
