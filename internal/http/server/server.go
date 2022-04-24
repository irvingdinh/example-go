package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/irvingdinh/example-go/internal/config"
	"github.com/irvingdinh/example-go/internal/http/handler"
)

type Server interface {
	Start() error
}

func New(handler handler.Handler) Server {
	server := &serverImpl{
		handler: handler,
	}

	server.withRouter()

	return server
}

type serverImpl struct {
	handler handler.Handler

	router *gin.Engine
}

func (i *serverImpl) Start() error {
	port := config.GetHTTPConfig().Port
	return i.router.Run(fmt.Sprintf(":%d", port))
}

func (i *serverImpl) withRouter() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	router.GET("/api/v1/tasks", i.handler.TaskHandler().GetTasks)
	router.POST("/api/v1/tasks", i.handler.TaskHandler().CreateTask)
	router.PATCH("/api/v1/tasks/:id/completed", i.handler.TaskHandler().MarkTaskAsCompleted)
	router.DELETE("/api/v1/tasks/:id", i.handler.TaskHandler().DeleteTask)

	i.router = router
}
