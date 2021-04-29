package model

type Article struct {
	*Model
	Title         string `bson:"title"`
	Desc          string `bson:"desc"`
	Content       string `bson:"Content"`
	CoverImageUrl string `bson:"Cover_image_url"`
	State         int    `bson:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
