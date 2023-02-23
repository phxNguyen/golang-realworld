package uploadmodel

import "golang-realworld/common"

const EntityName = "Upload"

type Upload struct {
	common.BaseModel `json:",inline"`
	common.Image     `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}
