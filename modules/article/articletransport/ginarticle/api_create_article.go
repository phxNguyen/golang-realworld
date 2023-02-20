package ginarticle

import (
	"github.com/gin-gonic/gin"
	"golang-cookie-blog/component"
	"golang-cookie-blog/modules/article/articlemodel"
	"golang-cookie-blog/modules/article/articleservice"
	"golang-cookie-blog/modules/article/articlestore"
	"net/http"
)

func CreateRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data articlemodel.ArticleCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := articlestore.NewSQLStore(appCtx.GetMainDbConnection())
		service := articleservice.NewCreateArticleService(store)

		if err := service.CreateArticle(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, data)
	}
}
