package controllers

import (
	"github.com/ahmetberke/wooker-api/internal/app/v1/errorss"
	"github.com/ahmetberke/wooker-api/internal/app/v1/response"
	"github.com/ahmetberke/wooker-api/internal/auth"
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type WordController struct {
	Service *service.WordService
	Auth *auth.Google
}

func (w *WordController) Get(c *gin.Context)  {

	var resp response.WordResponse

	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.InvalidWordID
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	id := uint(idInt)
	word, err := w.Service.FindByID(id)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.WordNotFound
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Word = models.ToWordDTO(word)

	c.JSON(resp.Code, resp)

}

func (w *WordController) New(c *gin.Context)  {

	var resp response.WordResponse

	userI, isExists := c.Get("x-user")
	if !isExists {
		resp.Code = http.StatusUnauthorized
		resp.Error = errorss.Unauthorized
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}
	loggedUser := userI.(*models.User)

	var wordDTO *models.WordDTO
	err := c.BindJSON(&wordDTO)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.CannotBind
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	word := models.ToWord(wordDTO)
	word.UserID = loggedUser.ID
	word.LanguageID = wordDTO.Language.ID

	wordS, err := w.Service.Save(word)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.UnableWordSaveToDB
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Word = models.ToWordDTO(wordS)

	c.JSON(resp.Code, resp)

}