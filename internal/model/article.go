package model

import "github.com/SuperArnold/GO_Blog/pkg/app"

type Article struct {
	*Model
	Title         string `bson:"title"`
	Desc          string `bson:"desc"`
	Content       string `bson:"Content"`
	CoverImageUrl string `bson:"Cover_image_url"`
	State         int    `bson:"state"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}
