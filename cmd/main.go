package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Article struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

func (Article) TableName() string {
	return "articles"
}

type ArticleUpdate struct {
	Title   *string `json:"title" gorm:"column:title;"`
	Content *string `json:"content" gorm:"column:content;"`
}

func (ArticleUpdate) TableName() string {
	return Article{}.TableName()
}
func main() {

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file ")
	}

	dsn := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to DB", err)
	}

	if err = runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	articles := router.Group("/articles")
	{
		articles.POST("/", func(c *gin.Context) {
			var article Article
			if err := c.ShouldBind(&article); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err,
				})
				return
			}

			if err := db.Create(&article).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, article)

		})

		articles.GET("/:id", func(c *gin.Context) {

			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err,
				})
				return
			}

			var article Article

			if err := db.First(&article, id).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err,
				})
				return
			}

			c.JSON(http.StatusOK, article)
		})

		articles.PATCH("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err,
				})
				return
			}

			var articleUpdate ArticleUpdate
			if err := c.ShouldBind(&articleUpdate); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
				return
			}

			if err := db.Table(Article{}.TableName()).Where("id = ?", id).Updates(&articleUpdate).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, articleUpdate)

		})

		articles.GET("/", func(c *gin.Context) {
			var listArticles []Article

			type Filter struct {
				HasLiked int `json:"has_liked" form:"has_liked"`
			}

			var filter Filter

			c.ShouldBind(&filter)

			dbFilter := db
			if filter.HasLiked > 0 {
				dbFilter = db.Where("has_liked = ?", filter.HasLiked)
			}

			if err := dbFilter.Find(&listArticles).Error; err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, listArticles)
		})

		articles.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
				return
			}

			if err := db.Table(Article{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
				return

			}

			c.JSON(http.StatusOK, gin.H{"ok": 1})
		})

	}

	return router.Run()
}
