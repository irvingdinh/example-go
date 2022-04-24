package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/irvingdinh/example-go/dto"
	"github.com/irvingdinh/example-go/internal/service/repository"
)

//go:generate mockery --name=Handler --case=snake
type Handler interface {
	TaskHandler() TaskHandler
}

func New(
	repository repository.Repository,
) Handler {
	return &handlerImpl{
		taskHandler: NewTaskHandler(repository),
	}
}

type handlerImpl struct {
	taskHandler TaskHandler
}

func (i *handlerImpl) TaskHandler() TaskHandler {
	return i.taskHandler
}

func abortWithError(c *gin.Context, code int, args ...error) {
	err := errors.New(strconv.Itoa(code))
	if len(args) == 1 {
		err = args[0]
	}

	c.JSON(code, &dto.ErrorResponse{Error: err.Error()})
}
