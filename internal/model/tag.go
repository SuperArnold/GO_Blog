package model

type Tag struct {
	*Model
	Name  string `bson:"name"`
	State int    `bson:"State"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
