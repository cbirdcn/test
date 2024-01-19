package main

import "github.com/gin-gonic/gin"

type LoginForm struct {
	User	string	`form:"user" binding:"required"`
	Password	string	`form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context){
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "logged in", "id": id, "page": page})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	r.Run(":8080")
}

//curl -X POST "http://127.0.0.1:8080/login?id=123&page=1" -d "user=user&password=password"
