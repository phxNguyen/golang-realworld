package articlemodel

type Filter struct {
	HasLiked int `json:"has_liked,omitempty" form:"has_liked" "`
}
