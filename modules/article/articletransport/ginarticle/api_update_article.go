package ginarticle

import (
	"github.com/gin-gonic/gin"
	"golang-realworld/modules/article/articlemodel"
	"golang-realworld/modules/article/articleservice"
	"golang-realworld/modules/article/articlestore"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateArticleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		var data articlemodel.ArticleUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := articlestore.NewArticleStore(db)
		service := articleservice.NewUpdateArticleService(store)

		if err := service.UpdateArticleService(c.Request.Context(), &data, id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, &data)

	}
}
