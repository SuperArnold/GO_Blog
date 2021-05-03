package model

import "github.com/SuperArnold/GO_Blog/pkg/app"

type Tag struct {
	*Model
	Name  string `bson:"name"`
	State int    `bson:"State"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
