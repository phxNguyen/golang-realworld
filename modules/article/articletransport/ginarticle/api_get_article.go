package ginarticle

import (
	"github.com/gin-gonic/gin"
	"golang-cookie-blog/modules/article/articleservice"
	"golang-cookie-blog/modules/article/articlestore"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetArticleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		store := articlestore.NewSQLStore(db)

		service := articleservice.NewGetArticle(store)

		data, err := service.FindArticleById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, data)

	}
}
