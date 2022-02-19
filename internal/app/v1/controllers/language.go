package controllers

import (
	"github.com/ahmetberke/wooker-api/internal/app/v1/errorss"
	"github.com/ahmetberke/wooker-api/internal/app/v1/response"
	"github.com/ahmetberke/wooker-api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LanguageController struct {
	Service *service.LanguageService
}

func NewLanguageController(service *service.LanguageService) *LanguageController {
	return &LanguageController{
		Service: service,
	}
}

func (l *LanguageController) All(c *gin.Context)  {
	var resp response.LanguagesResponse

	search := c.Query("search")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	languages, err := l.Service.GetAll(limit, search)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.SomethingIsWrong
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	resp.Languages = languages
	resp.Code = http.StatusOK
	c.JSON(resp.Code, resp)

}
