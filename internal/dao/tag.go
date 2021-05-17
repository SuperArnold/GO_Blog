package dao

import (
	"github.com/SuperArnold/GO_Blog/internal/model"
)

func (d *Dao) CreateTag(name string, state int, createdBy string) error {

	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id int, name string, state int, modifiedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{ID: id, ModifiedBy: modifiedBy},
	}

	return tag.Update(d.engine)
}

func (d *Dao) DeleteTag(id int) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	return tag.Delete(d.engine)
}
