package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang-realworld/component"
	"golang-realworld/component/uploadprovider"
	"golang-realworld/middleware"
	"golang-realworld/modules/article/articletransport/ginarticle"
	"golang-realworld/modules/upload/uploadtransport/ginupload"
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
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	dsn := os.Getenv("DbConnection")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to DB", err)
	}

	if err = runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {
	router := gin.Default()
	appCtx := component.NewAppContext(db, upProvider)
	router.Use(middleware.Recover(appCtx))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/upload", ginupload.Upload(appCtx))

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
