package router

import (
	"basicApi/utils"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"net/http"
	"strings"
)

type User struct {
	Username 	string `form:"username" json:"username" xml:"username" binding:"required"`
	AuthType 	string `form:"authType" json:"authType" xml:"authType" binding:"required"`
}
var (
	ValidAuthentications = []string{"user", "admin", "subscriber"}
)

func FindPostHandler(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{"data": articlecontroller.FindArticle()})
}

func InsertArticleHandler(c *gin.Context) {
	c.JSONP(http.StatusOK, articlecontroller.SaveArticle(c))
}

func loginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)

	if strings.Trim(user.Username, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username can't be empty"})
		return
	}
	if ! funk.ContainsString(ValidAuthentications, user.AuthType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid auth type"})
		return
	}

	// Note: This is just an example
	session.Set("user", user.Username)
	session.Set("authType", user.AuthType)

	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate session token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "authentication successful", "token": utils.GenerateToken(user.Username)})
}

func logoutHandler(c *gin.Context) {
	session := sessions.Default(c)

	// this would only be hit if the user was authenticated
	session.Delete("user")
	session.Delete("authType")
}