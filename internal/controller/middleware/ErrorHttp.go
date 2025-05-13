package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ErrorHttp(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("panic: %v\n", r)
			debug.PrintStack()
			c.Status(http.StatusInternalServerError)
		}
	}()
	c.Next()
}
