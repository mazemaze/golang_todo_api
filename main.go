package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/mazemaze/golang_todo/controller"
)

type SessionInfo struct {
	name interface{}
}

var LoginInfo SessionInfo

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3000/list"}
	config.AllowWildcard = true
	config.AllowCredentials = true
	r.Use(cors.New(config))

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.POST("/user", controller.SaveSession)
	r.GET("/user", controller.GetSession)
	r.POST("/task", controller.AddTask)

	r.GET("/task/:id", controller.GetTaskByID)

	r.GET("/task", controller.GetTasks)

	r.Run()
}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		name := session.Get("name")
		if name == nil {
			log.Println("User name is not exist")
			c.Redirect(http.StatusMovedPermanently, "/")
			c.Abort()
		} else {
			session.Set("Name", "Guest User")
			c.Next()
		}
		log.Println("Checking session is over")
	}
}
