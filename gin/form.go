package main

import "github.com/gin-gonic/gin"

type LoginForm struct {
	User	string	`form:"user" binding:"required"`
	Password	string	`form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context){
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	r.Run(":8080")
}

// curl -X POST http://127.0.0.1:8080/login -d "user=user&password=password"
