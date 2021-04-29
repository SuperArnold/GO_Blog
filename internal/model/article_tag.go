package model

type ArticleTag struct {
	*Model
	TagID     int `bson:"tag_id"`
	ArticleID int `bson:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
