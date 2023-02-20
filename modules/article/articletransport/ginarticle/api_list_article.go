package ginarticle

import (
	"github.com/gin-gonic/gin"
	"golang-cookie-blog/common"
	"golang-cookie-blog/modules/article/articlemodel"
	"golang-cookie-blog/modules/article/articleservice"
	"golang-cookie-blog/modules/article/articlestore"
	"gorm.io/gorm"
	"net/http"
)

func ListArticleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var filter articlemodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		_ = paging.Validate()

		store := articlestore.NewSQLStore(db)
		service := articleservice.NewListArticleService(store)

		data, err := service.ListArticleService(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{"paging": paging, "data": data, "filter": filter})
	}
}
