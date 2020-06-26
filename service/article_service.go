package service

import "basicApi/schema"

type ArticleService interface {
	SaveArticle(articlesSchema schema.ArticlesSchema) schema.ArticlesSchema
	FindArticle() []schema.ArticlesSchema
}

type articlesService struct {
	articles []schema.ArticlesSchema
}

func New() ArticleService {
	return &articlesService{
		articles: []schema.ArticlesSchema{},
	}
}

func (service *articlesService) SaveArticle(obj schema.ArticlesSchema) schema.ArticlesSchema {
	service.articles = append(service.articles, obj)
	return obj
}

func (service *articlesService) FindArticle() []schema.ArticlesSchema {
	return service.articles
}