package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/irvingdinh/example-go/internal/component/logger"
	"github.com/irvingdinh/example-go/internal/model"
)

type (
	CreateTaskArgs struct {
		Summary string
	}
)

//go:generate mockery --name=TaskRepository --case=snake
type TaskRepository interface {
	CreateTask(ctx context.Context, args CreateTaskArgs) (*model.Task, error)
	GetTasks(ctx context.Context) ([]*model.Task, error)
	MarkTaskAsCompleted(ctx context.Context, id uint64) error
	DeleteTask(ctx context.Context, id uint64) error
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepositoryImpl{
		db: db,
	}
}

type taskRepositoryImpl struct {
	db *gorm.DB
}

func (i *taskRepositoryImpl) CreateTask(ctx context.Context, args CreateTaskArgs) (*model.Task, error) {
	log := logger.CToL(ctx, "CreateTask")
	log.WithField("args", args)

	task := &model.Task{
		Summary: args.Summary,
	}

	err := i.db.WithContext(ctx).Create(&task).Error
	if err != nil {
		log.WithError(err).Errorf("gorm.DB returns error when .Create: %s", err.Error())
		return nil, err
	}

	return task, nil
}

func (i *taskRepositoryImpl) GetTasks(ctx context.Context) ([]*model.Task, error) {
	log := logger.CToL(ctx, "GetTasks")

	var tasks []*model.Task

	err := i.db.WithContext(ctx).Find(&tasks).Error
	if err != nil {
		log.WithError(err).Errorf("gorm.DB returns error when .Find: %s", err.Error())
		return nil, err
	}

	return tasks, nil
}

func (i *taskRepositoryImpl) MarkTaskAsCompleted(ctx context.Context, id uint64) error {
	log := logger.CToL(ctx, "MarkTaskAsCompleted")
	log = log.WithField("id", id)

	var task *model.Task

	err := i.db.WithContext(ctx).Where("id = ?", id).First(&task).Error
	if err != nil {
		log.WithError(err).Errorf("gorm.DB returns error when .First: %s", err.Error())
		return err
	}

	task.IsCompleted = true

	err = i.db.WithContext(ctx).Save(task).Error
	if err != nil {
		log.WithError(err).Errorf("gorm.DB returns error when .Save: %s", err.Error())
		return err
	}

	return nil
}

func (i *taskRepositoryImpl) DeleteTask(ctx context.Context, id uint64) error {
	log := logger.CToL(ctx, "DeleteTask")
	log = log.WithField("id", id)

	err := i.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Task{}).Error
	if err != nil {
		log.WithError(err).Errorf("gorm.DB returns error when .Delete: %s", err.Error())
		return err
	}

	return nil
}
