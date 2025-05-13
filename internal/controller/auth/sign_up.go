package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"glossika/internal/domain/er"
	authucase "glossika/internal/usecase/auth"
	"glossika/internal/util"
)

func (h *Handler) SignUp(c *gin.Context) {
	var req authucase.SignUpIn
	if !util.ParseValidate(c, &req) {
		return
	}

	if err := h.Usecase.SignUp(c.Request.Context(), req); err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.Status(http.StatusOK)
}
