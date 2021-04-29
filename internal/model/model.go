package model

import (
	"fmt"

	"github.com/SuperArnold/GO_Blog/global"
	"github.com/SuperArnold/GO_Blog/pkg/setting"
	"github.com/globalsign/mgo/bson"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedBy  string        `bson:"created_by"`
	ModifiedBY string        `bson:"modified_bY"`
	CreateOn   string        `bson:"create_on"`
	ModifiedOn string        `bson:"modified_on"`
	DeletedOn  string        `bson:"deleted_on"`
	IsDel      int           `bson:"is_del"`
}

func NewDBEngine(DatabaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DatabaseSetting.Host,
		DatabaseSetting.Port,
		DatabaseSetting.User,
		DatabaseSetting.DB,
		DatabaseSetting.Password,
	)
	db, err := gorm.Open("postgres", s)
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	// db.DB().SetMaxIdleConns(DatabaseSetting.)

	return db, nil
}
