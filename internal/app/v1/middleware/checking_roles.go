package middleware

import (
	"github.com/ahmetberke/wooker-api/internal/app/v1/errorss"
	"github.com/ahmetberke/wooker-api/internal/app/v1/response"
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *Manager) IsAdminChecking(c *gin.Context) {

	var resp response.UsersResponse

	userI, ok := c.Get("x-user")
	if !ok {
		resp.Code = http.StatusUnauthorized
		resp.Error = errorss.Unauthorized
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	user := userI.(*models.User)
	if !user.IsAdmin {
		resp.Code = http.StatusUnauthorized
		resp.Error = errorss.Unauthorized
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	c.Next()

}

func (m *Manager) IsLoggedUserOrAdminChecking(c *gin.Context) {
	var resp response.UsersResponse

	username := c.Param("username")
	userI, ok := c.Get("x-user")
	if !ok {
		resp.Code = http.StatusUnauthorized
		resp.Error = errorss.Unauthorized
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	user := userI.(*models.User)
	if !user.IsAdmin && (username != user.Username) {
		resp.Code = http.StatusUnauthorized
		resp.Error = errorss.Unauthorized
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	c.Next()

}