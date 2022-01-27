package controller

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mazemaze/golang_todo/models"
)

type Name struct {
	Name string `json:"name"`
}

func AddTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	models.Tasks = append(models.Tasks, newTask)
	c.IndentedJSON(http.StatusOK, newTask)
}

func GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range models.Tasks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func SaveSession(c *gin.Context) {
	var name Name
	if err := c.BindJSON(&name); err != nil {
		return
	}
	session := sessions.Default(c)
	session.Set("name", name.Name)
	session.Save()
	v := session.Get("name")
	log.Println(v)
	c.IndentedJSON(200, gin.H{"name": v})
}

func GetSession(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("name")
	if v == nil {
		log.Println("User Not Found")
		return
	}
	c.IndentedJSON(http.StatusOK, v)
}
