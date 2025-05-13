package er

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Bind(c *gin.Context, err error) {
	var appErr *AppError
	if !errors.As(err, &appErr) {
		c.JSON(
			http.StatusInternalServerError,
			NewAPIErr(
				"",
				Unknown.String(),
				fmt.Sprintf("bind app error: original error: %v", err.Error()),
			),
		)
		return
	}

	c.JSON(
		appErr.HTTPCode(),
		NewAPIErr(
			appErr.traceID.String(),
			appErr.code.String(),
			CodeToMsg[appErr.code],
		),
	)
}
