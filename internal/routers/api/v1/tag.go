package v1

import (
	"github.com/SuperArnold/GO_Blog/pkg/app"
	"github.com/SuperArnold/GO_Blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}
type Article struct{}

func NewTag() Tag {
	return Tag{}
}

func NewArticle() Article {
	return Article{}
}

func (t Tag) Get(c *gin.Context) {}

// @Summary 取得多個標籤
// @Produce json
// @Param name query string false "標籤名稱" maxlength(100)
// @Param state query int false "狀態" Enums(0,1) default(1)
// @Param page query int false "頁碼"
// @Param page_size query int false "每頁數量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {}

// @Summary 新增標籤
// @Produce json
// @Param name query string false "標籤名稱" minlength(3) maxlength(100)
// @Param state query int false "狀態" Enums(0,1) default(1)
// @Param create_by body string true "建立者" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}

// @Summary 更新標籤
// @Produce json
// @Param id path int true "標籤ID"
// @Param name body string false "標籤名稱" minlength(3) maxlength(100)
// @Param state body int false "狀態" Enums(0,1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {}

// @Summary 刪除標籤
// @Produce json
// @Param id path int true "標籤ID"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}

func (t Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (t Article) List(c *gin.Context)   {}
func (t Article) Create(c *gin.Context) {}
func (t Article) Update(c *gin.Context) {}
func (t Article) Delete(c *gin.Context) {}
