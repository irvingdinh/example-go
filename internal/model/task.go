package model

type Task struct {
	Model

	Summary     string
	IsCompleted bool
}

func (i Task) AsMap() M {
	return M{
		"id":           i.ID,
		"summary":      i.Summary,
		"is_completed": i.IsCompleted,
		"created_at":   timeAsMapValue(i.CreatedAt),
		"updated_at":   timeAsMapValue(i.UpdatedAt),
		"deleted_at":   nullTimeAsMapValue(i.DeletedAt),
	}
}
