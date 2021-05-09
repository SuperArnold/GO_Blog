package model

import (
	"github.com/SuperArnold/GO_Blog/global"
	"github.com/SuperArnold/GO_Blog/pkg/app"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name  string `bson:"name"`
	State int    `bson:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}
	db = db.Where("state=?", t.State)

	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Where(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}

	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	global.Logger.Info("table is exist")
	global.Logger.Infof("table is exist : %s", t.TableName())
	has := db.HasTable(t)
	if !has {
		global.Logger.Info("table not exist")
	}

	return db.Create(&t).Error
}

// func (t Tag) Update(db *gorm.DB) error {
// 	db = db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID, 0)
// 	return db.Update(t).Error
// }

// func (t Tag) Delete(db *gorm.DB) error {
// 	db = db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID, 0)
// 	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
// }
