package ginarticle

import (
	"github.com/gin-gonic/gin"
	"golang-realworld/common"
	"golang-realworld/modules/article/articleservice"
	"golang-realworld/modules/article/articlestore"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetArticleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := articlestore.NewArticleStore(db)

		service := articleservice.NewGetArticle(store)

		data, err := service.FindArticleById(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, data)

	}
}
