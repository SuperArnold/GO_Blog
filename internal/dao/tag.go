package dao

import (
	"github.com/SuperArnold/GO_Blog/global"
	"github.com/SuperArnold/GO_Blog/internal/model"
)

func (d *Dao) CreateTag(name string, state int, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	global.Logger.Infof("CCCCCC : %v", tag.CreatedBy)

	return tag.Create(d.engine)
}
