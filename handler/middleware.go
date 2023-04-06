package handler

import (
	"net/http"
	"paa/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) IsLogin(c *gin.Context) {
	cookie, err := c.Request.Cookie("jwt")
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	// Extract token string
	tokenString := cookie.Value

	// Validate token
	_, err = utils.ValidateToken(tokenString)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.Next()
}
