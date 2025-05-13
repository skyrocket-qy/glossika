package authcontroller

import (
	"net/http"

	"glossika/internal/domain/er"
	authucase "glossika/internal/usecase/auth"
	"glossika/internal/util"

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
