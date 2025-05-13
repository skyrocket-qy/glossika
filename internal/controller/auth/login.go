package authcontroller

import (
	"net/http"

	"glossika/internal/domain/er"
	authucase "glossika/internal/usecase/auth"
	"glossika/internal/util"

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
