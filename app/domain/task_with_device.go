package domain

import "time"

// ListByDeviceUUID returns nested tasks selected by deviceUUID.
func (t *Task) ListByDeviceUUID(deviceUUID string) (*[]Task, error) {
	var tasks []Task
	treeableTasks, err := t.taskAccessor.ListTreeByDeviceUUID(deviceUUID)
	if err != nil {
		return &tasks, err
	}

	for _, treeableTask := range *treeableTasks {
		task := Task{
			ID:        treeableTask.ID,
			Title:     treeableTask.Title,
			Type:      TaskType(treeableTask.Type).String(),
			CreatedAt: treeableTask.CreatedAt.In(time.Local).Format("2006-01-02 15:04:05"),
			UpdatedAt: treeableTask.UpdatedAt.In(time.Local).Format("2006-01-02 15:04:05"),
			Depth:     treeableTask.Depth,
		}
		if treeableTask.CompletedAt != nil {
			task.CompletedAt = treeableTask.CompletedAt.In(time.Local).Format("2006-01-02 15:04:05")
		}
		if treeableTask.StartsAt != nil {
			task.StartsAt = treeableTask.StartsAt.In(time.Local).Format("2006-01-02 15:04:05")
		}
		if treeableTask.ExpiresAt != nil {
			task.ExpiresAt = treeableTask.ExpiresAt.In(time.Local).Format("2006-01-02 15:04:05")
		}
		task = task.mapChildren(treeableTask.Children)
		tasks = append(tasks, task)
	}

	return &tasks, nil
}

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
		CreatedAt: (*repositoryTask).CreatedAt.In(time.Local).Format("2006-01-02 15:04:05"),
		UpdatedAt: (*repositoryTask).UpdatedAt.In(time.Local).Format("2006-01-02 15:04:05"),
		Depth:     1,
	}

	_, err = t.taskAccessor.CreateDeviceTask(deviceUUID, task.ID)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
