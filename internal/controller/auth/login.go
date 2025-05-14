package authcontroller

import (
	"net/http"

	"recsvc/internal/domain/er"
	authucase "recsvc/internal/usecase/auth"
	"recsvc/internal/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var req authucase.LoginIn
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	out, err := h.Usecase.Login(c.Request.Context(), req)
	if err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.JSON(http.StatusOK, out)
}
