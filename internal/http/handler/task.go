package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/irvingdinh/example-go/dto"
	"github.com/irvingdinh/example-go/internal/component/logger"
	"github.com/irvingdinh/example-go/internal/model/converter"
	"github.com/irvingdinh/example-go/internal/service/repository"
)

//go:generate mockery --name=TaskHandler --case=snake
type TaskHandler interface {
	CreateTask(c *gin.Context)
	GetTasks(c *gin.Context)
	MarkTaskAsCompleted(c *gin.Context)
	DeleteTask(c *gin.Context)
}

func NewTaskHandler(
	repository repository.Repository,
) TaskHandler {
	return &taskHandlerImpl{
		repository: repository,
	}
}

type taskHandlerImpl struct {
	repository repository.Repository
}

func (i *taskHandlerImpl) CreateTask(c *gin.Context) {
	req, err := i.bindCreateTaskRequest(c)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	log := logger.CToL(c.Request.Context(), "CreateTask")
	log = log.WithField("req", req)

	task, err := i.repository.TaskRepository().CreateTask(
		logger.LToC(c.Request.Context(), log),
		repository.CreateTaskArgs{
			Summary: req.Summary,
		},
	)
	if err != nil {
		log.WithField("err", err).Errorf("CreateTask returns error: %s", err.Error())
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &dto.CreateTaskResponse{
		Data: converter.TaskToDTO(task),
	})
}

func (i *taskHandlerImpl) GetTasks(c *gin.Context) {
	log := logger.CToL(c.Request.Context(), "GetTasks")

	tasks, err := i.repository.TaskRepository().GetTasks(
		logger.LToC(c.Request.Context(), log),
	)
	if err != nil {
		log.WithField("err", err).Errorf("GetTasks returns error: %s", err.Error())
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &dto.GetTasksResponse{
		Data: converter.TasksToDTO(tasks),
	})
}

func (i *taskHandlerImpl) MarkTaskAsCompleted(c *gin.Context) {
	req, err := i.bindMarkTaskAsCompletedRequest(c)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	log := logger.CToL(c.Request.Context(), "MarkTaskAsCompleted")
	log = log.WithField("req", req)

	err = i.repository.TaskRepository().MarkTaskAsCompleted(
		logger.LToC(c.Request.Context(), log),
		req.ID,
	)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			abortWithError(c, http.StatusNotFound)
			return
		}

		log.WithField("err", err).Errorf("MarkTaskAsCompleted returns error: %s", err.Error())
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func (i *taskHandlerImpl) DeleteTask(c *gin.Context) {
	req, err := i.bindDeleteTaskRequest(c)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	log := logger.CToL(c.Request.Context(), "DeleteTask")
	log = log.WithField("req", req)

	err = i.repository.TaskRepository().DeleteTask(
		logger.LToC(c.Request.Context(), log),
		req.ID,
	)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			abortWithError(c, http.StatusNotFound)
			return
		}

		log.WithField("err", err).Errorf("DeleteTask returns error: %s", err.Error())
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func (i *taskHandlerImpl) bindCreateTaskRequest(c *gin.Context) (*dto.CreateTaskRequest, error) {
	var req *dto.CreateTaskRequest

	if err := c.ShouldBind(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func (i *taskHandlerImpl) bindMarkTaskAsCompletedRequest(c *gin.Context) (*dto.MarkTaskAsCompletedRequest, error) {
	var req *dto.MarkTaskAsCompletedRequest

	if err := c.ShouldBindUri(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func (i *taskHandlerImpl) bindDeleteTaskRequest(c *gin.Context) (*dto.DeleteTaskRequest, error) {
	var req *dto.DeleteTaskRequest

	if err := c.ShouldBindUri(&req); err != nil {
		return nil, err
	}

	return req, nil
}
