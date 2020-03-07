package domain

// CreateWithDevice saves task and device_task record through persistence layer.
func (t *Task) CreateWithDevice(deviceUUID string, title string) (*Task, error) {
	repositoryTask, err := t.taskAccessor.Create(title)
	if err != nil {
		return nil, err
	}

	task := Task{
		ID:        (*repositoryTask).ID,
		Title:     (*repositoryTask).Title,
		Type:      TaskType((*repositoryTask).Type).String(),
		CreatedAt: (*repositoryTask).CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: (*repositoryTask).UpdatedAt.Format("2006-01-02 15:04:05"),
		Depth:     1,
	}

	_, err = t.taskAccessor.CreateDeviceTask(deviceUUID, task.ID)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
