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

func DeleteArticleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		store := articlestore.NewSQLStore(db)
		service := articleservice.NewDeleteArticleService(store)

		if err := service.DeleteArticleService(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, common.NewShortSuccessResponse(true))

	}
}
