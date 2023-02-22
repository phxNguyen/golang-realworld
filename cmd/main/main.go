package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang-realworld/component"
	"golang-realworld/middleware"
	"golang-realworld/modules/article/articletransport/ginarticle"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Article struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

func (Article) TableName() string {
	return "articles"
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
	appCtx := component.NewAppContext(db)
	router.Use(middleware.Recover(appCtx))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	articles := router.Group("/articles")
	{
		articles.POST("/", ginarticle.CreateRestaurantHandler(appCtx))

		articles.GET("/:id", ginarticle.GetArticleHandler(db))

		articles.PATCH("/:id", ginarticle.UpdateArticleHandler(db))

		articles.GET("/", ginarticle.ListArticleHandler(db))

		articles.DELETE("/:id", ginarticle.DeleteArticleHandler(db))

	}

	return router.Run()
}
