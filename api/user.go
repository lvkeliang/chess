package api

import (
	"Gone/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

// 注册
func Register(h *model.Hub, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		//获取参数
		username := c.PostForm("username")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

		//数据验证
		if len(username) == 0 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "用户名不能为空",
			})
			return
		}
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "手机号必须为11位",
			})
			return
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "密码不能少于6位",
			})
			return
		}

		//判断手机号是否存在
		var user model.User
		db.Where("telephone = ?", telephone).First(&user)
		if user.ID != 0 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "用户已存在",
			})
			return
		}

		//创建用户
		hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    500,
				"message": "密码加密错误",
			})
			return
		}
		newUser := model.User{
			Username:  username,
			Telephone: telephone,
			Password:  string(hasedPassword),
		}
		db.Create(&newUser)

		//返回结果
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "注册成功",
		})
	}
}

// 登录
func Login(h *model.Hub, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		//获取参数
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

		//数据验证
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "手机号必须为11位",
			})
			return
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "密码不能少于6位",
			})
			return
		}

		//判断手机号是否存在
		var user model.User
		db.Where("telephone = ?", telephone).First(&user)
		if user.ID == 0 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "用户不存在",
			})
			return
		}

		//判断密码是否正确
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "密码错误",
			})
		}

		//设置token并返回结果
		/*c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "登录成功",
		})*/
		SetToken(user, c)
	}
}
