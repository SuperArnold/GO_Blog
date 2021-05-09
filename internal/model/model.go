package model

import (
	"fmt"
	"time"

	"github.com/SuperArnold/GO_Blog/global"
	"github.com/SuperArnold/GO_Blog/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Model struct {
	// ID         int    `bson:"_id"`
	CreatedBy  string `bson:"created_by"`
	ModifiedBY string `bson:"modified_bY"`
	CreateOn   string `bson:"create_on"`
	ModifiedOn string `bson:"modified_on"`
	DeletedOn  string `bson:"deleted_on"`
	IsDel      int    `bson:"is_del"`
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

	global.Logger.Infof("FFFFFFF: %v", DatabaseSetting.Host)

	db, err := gorm.Open("postgres", s)
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	// db.DB().SetMaxIdleConns(DatabaseSetting.)

	global.Logger.Info("DB connext GOOD")
	return db, nil
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}
