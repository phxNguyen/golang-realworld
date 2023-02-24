package component

import (
	"golang-realworld/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appContext struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider) *appContext {
	return &appContext{
		db:         db,
		upProvider: upProvider,
	}
}

func (ctx *appContext) GetMainDbConnection() *gorm.DB {

	return ctx.db
}

func (ctx *appContext) UploadProvider() uploadprovider.UploadProvider {

	return ctx.upProvider
}
