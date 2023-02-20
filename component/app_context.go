package component

import "gorm.io/gorm"

type AppContext interface {
	GetMainDbConnection() *gorm.DB
}

type appContext struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{db: db}
}

func (ctx *appContext) GetMainDbConnection() *gorm.DB {

	return ctx.db
}
