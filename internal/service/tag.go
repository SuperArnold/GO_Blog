package service

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State int    `form:"state,default=1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State int    `form:"state,default=1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     int    `form:"state,default=1"`
}

type UpdateTagRequest struct {
	ID         int    `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"required,min=3,max=100"`
	State      int    `form:"state,default=01"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID int `form:"id" binding:"required,gte=1"`
}
