package controller

import (
	"basicApi/schema"
	"basicApi/service"
	"github.com/gin-gonic/gin"
)

type ArticleController interface {
	FindArticle() []schema.ArticlesSchema
	SaveArticle(ctx *gin.Context) schema.ArticlesSchema
}

type controllers struct {
	service service.ArticleService
}

func (c controllers) FindArticle() []schema.ArticlesSchema {
	return c.service.FindArticle()
}

func (c controllers) SaveArticle(ctx *gin.Context) schema.ArticlesSchema {
	var article schema.ArticlesSchema
	ctx.BindJSON(&article)
	c.service.SaveArticle(article)
	return article
}

func New(service service.ArticleService) *controllers {
	return &controllers{
		service: service,
	}
}