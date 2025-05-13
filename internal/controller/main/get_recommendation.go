package maincontroller

import (
	"net/http"

	"glossika/internal/domain/er"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRecommendation(c *gin.Context) {
	out, err := h.Usecase.GetRecommendation(c.Request.Context())
	if err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.JSON(http.StatusOK, out)
}
