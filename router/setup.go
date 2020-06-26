package router

import (
	"basicApi/controller"
	"basicApi/middlewares"
	"basicApi/service"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	articleService service.ArticleService = service.New()
	articlecontroller controller.ArticleController = controller.New(articleService)
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	store := sessions.NewCookieStore([]byte("ABCSECRET"))
	router.Use(gin.Recovery(), middlewares.Logger(), sessions.Sessions("sessionName", store))

	api := router.Group("users/v1")
	api.POST("/login", loginHandler)
	api.GET("/logout", logoutHandler)
	{
		articleRoute := api.Group("/")
		articleRoute.Use(middlewares.OpenAuth("abc", "abc123"), middlewares.AuthenticationRequired())
		{
			articleRoute.GET("/articles", FindPostHandler)
			articleRoute.POST("/articles", InsertArticleHandler)
		}
	}
	{
		basicAuth := api.Group("/")
		basicAuth.Use(middlewares.AuthenticationRequired("admin"))
		{
			basicAuth.GET("/info", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "hello world",
				})
			})
		}
	}
	return router
}