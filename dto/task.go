package dto

type Task struct {
	ID          uint64  `json:"id"`
	Summary     string  `json:"summary"`
	IsCompleted bool    `json:"is_completed"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at"`
}

type CreateTaskRequest struct {
	Summary string `json:"summary" binding:"min=3,required"`
}

type CreateTaskResponse struct {
	Data *Task `json:"data"`
}

type GetTasksResponse struct {
	Data []*Task `json:"data"`
}

type MarkTaskAsCompletedRequest struct {
	ID uint64 `uri:"id" binding:"min=1,required"`
}

type DeleteTaskRequest struct {
	ID uint64 `uri:"id" binding:"min=1,required"`
}
