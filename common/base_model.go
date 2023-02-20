package common

import "time"

type BaseModel struct {
	Id        int        `json:"id" gorm:"column:id"`
	Status    int        `json:"status" gorm:"column:status"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
