package controllers

import (
	"errors"
	"github.com/ahmetberke/wooker-api/internal/app/v1/errorss"
	"github.com/ahmetberke/wooker-api/internal/app/v1/response"
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type WordController struct {
	Service *service.WordService
}

func NewWordController(wordService *service.WordService) *WordController {
	return &WordController{
		Service: wordService,
	}
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

func (w *WordController) All(c *gin.Context)  {

	var resp response.WordsResponse

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	different := false
	if c.Query("different") == "true" {
		different = true
	}

	username := c.Query("username")
	languageCode := c.Query("language-id")

	words, err := w.Service.GetAll(limit, username, languageCode, different)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.CannotWordsRetrieved
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	resp.Code = http.StatusOK
	for _, w := range words {
		resp.Words = append(resp.Words, *models.ToWordDTO(&w))
	}

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

	wordS, err := w.Service.Save(word)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Code = http.StatusBadRequest
			resp.Error = errorss.LanguageNotFound
			c.AbortWithStatusJSON(resp.Code, resp)
			return
		}
		if errorss.IsAlreadyExistsErr(err) {
			resp.Code = http.StatusBadRequest
			resp.Error = errorss.WordAlreadyExists
			c.AbortWithStatusJSON(resp.Code, resp)
			return
		}
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.SomethingIsWrong
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Word = models.ToWordDTO(wordS)

	c.JSON(resp.Code, resp)

}

func (w *WordController) Delete(c *gin.Context) {
	var resp response.WordResponse

	userI, isExists := c.Get("x-user")
	if !isExists {
		resp.Code = http.StatusUnauthorized
		resp.Error = errorss.Unauthorized
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}
	loggedUser := userI.(*models.User)

	idI, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.InvalidWordID
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}
	if idI < 0 {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.InvalidWordID
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	id := uint(idI)

	err = w.Service.Delete(loggedUser.ID, id)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.WordNotFound
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	resp.Code = http.StatusOK
	c.JSON(resp.Code, resp)
	return

}

func (w *WordController) Update(c *gin.Context) {
	var resp response.WordResponse

	userI, isExists := c.Get("x-user")
	if !isExists {
		resp.Code = http.StatusUnauthorized
		resp.Error = errorss.Unauthorized
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}
	loggedUser := userI.(*models.User)

	idI, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.InvalidWordID
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}
	if idI < 0 {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.InvalidWordID
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	id := uint(idI)

	var wordDTO models.WordDTO
	err = c.ShouldBindJSON(&wordDTO)
	if err  != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.CannotBind
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	word := models.ToWord(&wordDTO)
	word.UserID = loggedUser.ID
	word.ID = id

	wordU, err := w.Service.Update(word)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.WordNotUpdate
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Word = models.ToWordDTO(wordU)
	c.JSON(resp.Code, resp)
	return

}