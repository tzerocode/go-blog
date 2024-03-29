package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog/pkg/app"
	"go-blog/pkg/errcode"
)

type Article struct {
}

func NewArticles() Article {
	return Article{}
}

func (t Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (t Article) List(c *gin.Context)   {}
func (t Article) Create(c *gin.Context) {}
func (t Article) Update(c *gin.Context) {}
func (t Article) Delete(c *gin.Context) {}
