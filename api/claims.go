package api

import (
	"Gone/model"
	util "Gone/uitl"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtKey = []byte("12190711")

func SetToken(user model.User, c *gin.Context) {

	// 创建 JWT
	expireTime := time.Now().Add(time.Hour * 24).Unix()
	claims := &model.Claims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"message": "Can't generate JWT"})
		util.RespUnexceptedError(c)
		return
	}
	// 返回 JWT
	//c.JSON(http.StatusOK, gin.H{"token": tokenString})
	c.SetCookie("token", tokenString, 604800, "", "/", false, false)
	util.RespSetTokenSuccess(c, tokenString)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 token
		//tokenString := c.GetHeader("Authorization")

		//尝试从Cookie中获取 token
		tokenString, err := c.Cookie("token")

		if tokenString == "" || err != nil {
			// 尝试从请求头中获取 token
			tokenString = c.GetHeader("Authorization")

			if tokenString == "" {
				util.RespDidNotLogin(c)
				c.Abort()
				return
			}
		}

		// 验证 token
		token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			util.RespInvalidToken(c)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			c.Set("ID", claims.ID)
		} else {
			util.RespInvalidToken(c)
			c.Abort()
			return
		}

		c.Next()
	}
}
