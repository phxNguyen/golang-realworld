package articlemodel

import (
	"errors"
	"golang-realworld/common"
	"strings"
)

const EntityName = "Article"

type Article struct {
	common.BaseModel `json:",inline"`
	Id               int    `json:"id" gorm:"column:id;"`
	AuthorId         int    `json:"author_id" gorm:"column:id;"`
	Title            string `json:"title" gorm:"column:title;"`
	Content          string `json:"content" gorm:"column:content;"`
	Logo             *common.Image
}

func (Article) TableName() string {
	return "articles"
}

type ArticleCreate struct {
	Title   string        `json:"title" gorm:"column:title;"`
	Content string        `json:"content" gorm:"column:content;"`
	Logo    *common.Image `json:"logo" gorm:"column:logo;"`
}

func (ArticleCreate) TableName() string {
	return "articles"
}

type ArticleUpdate struct {
	Title   *string `json:"title" gorm:"column:title;"`
	Content *string `json:"content" gorm:"column:content;"`
	Logo    *common.Image
}

func (ArticleUpdate) TableName() string {
	return Article{}.TableName()
}

// Validate model business
func (article *ArticleCreate) Validate() error {

	article.Title = strings.TrimSpace(article.Title)
	article.Content = strings.TrimSpace(article.Content)
	if len(article.Title) == 0 {
		return errors.New("article title cannot be blank")
	}
	if len(article.Content) == 0 {
		return errors.New("article content cannot be blank")
	}
	return nil
}
