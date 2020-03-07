package middleware

import (
	"log"
	"net/http"
	"netdisk/dao"
	"netdisk/help/huser"
	"netdisk/model"
	"netdisk/utils/ygjwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "权限不足"})
	}

	tokenString = tokenString[7:]
	token, claims, err := ygjwt.ParseToken(tokenString)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "权限不足"})
		c.Abort()
		return
	}

	db := dao.NewDB()
	var user model.UserModel
	if db.First(&user, "id=?", claims.UserId).RecordNotFound() {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "权限不足"})
		c.Abort()
		return
	}

	c.Set("user", huser.HideUserInfo(&user))
	tmp, _ := c.Get("user")
	log.Println(tmp)
	c.Next()
}
