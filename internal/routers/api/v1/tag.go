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

func (t Tag) Get(c *gin.Context)    {}
func (t Tag) List(c *gin.Context)   {}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}

func (t Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (t Article) List(c *gin.Context)   {}
func (t Article) Create(c *gin.Context) {}
func (t Article) Update(c *gin.Context) {}
func (t Article) Delete(c *gin.Context) {}
