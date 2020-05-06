package domain

import (
	"time"

	"github.com/whalepod/milelane/app/domain/repository"
)

// TaskType contains task type, used to decide which function is available to client.
type TaskType uint

const (
	// TypeTask is simple task, which can be done.
	TypeTask TaskType = iota * 10
	// TypeLane is abstract task like project, which can not be simply done.
	TypeLane
)

// TaskAccessor is interface explaining availble methods to approach persistence layer.
// Implementation(s) of TaskAccessor is/are
// - TaskRepository.
// - TaskAccessorMock (in test).
type TaskAccessor interface {
	ListTree() (*[]repository.TreeableTask, error)
	ListTreeByDeviceUUID(deviceUUID string) (*[]repository.TreeableTask, error)
	FindTreeByID(id uint) (*repository.TreeableTask, error)
	Create(title string) (*repository.Task, error)
	UpdateCompletedAt(id uint, completedAt time.Time) error
	UpdateStartsAt(id uint, startsAt *time.Time) error
	UpdateExpiresAt(id uint, expiresAt *time.Time) error
	UpdateTitle(id uint, title string) error
	UpdateType(id uint, taskType uint) error
	DeleteAncestorTaskRelations(taskID uint) error
	CreateTaskRelationsBetweenTasks(parentTaskID uint, childTaskID uint) error
	CreateDeviceTask(deviceUUID string, taskID uint) (*repository.DeviceTask, error)
}

// Task is struct for domain, not for gorm.
type Task struct {
	taskAccessor TaskAccessor
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	CompletedAt  string `json:"completed_at"`
	StartsAt     string `json:"starts_at"`
	ExpiresAt    string `json:"expires_at"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Depth        uint   `json:"depth"`
	Children     []Task `json:"children"`
}

// NewTask returns Task struct with TaskAccessor.
func NewTask(ta TaskAccessor) (*Task, error) {
	var t Task
	t.taskAccessor = ta

	return &t, nil
}

// List returns nested tasks.
func (t *Task) List() (*[]Task, error) {
	var tasks []Task
	treeableTasks, err := t.taskAccessor.ListTree()
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

// Find returns task and its descendants.
func (t *Task) Find(id uint) (*Task, error) {
	var task Task
	treeableTask, err := t.taskAccessor.FindTreeByID(id)
	if err != nil {
		return &task, err
	}

	task = Task{
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

	return &task, nil
}

func (t *Task) mapChildren(treeableTasks []repository.TreeableTask) Task {
	var children []Task
	for _, treeableTask := range treeableTasks {
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
		children = append(children, task)
	}
	(*t).Children = children

	return *t
}

// Create saves task record through persistence layer.
func (t *Task) Create(title string) (*Task, error) {
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

	return &task, nil
}

// Complete makes a task done.
func (t *Task) Complete(id uint) error {
	err := t.taskAccessor.UpdateCompletedAt(id, time.Now())
	if err != nil {
		return err
	}

	return nil
}

// UpdateTerm sets a task activation term.
func (t *Task) UpdateTerm(id uint, startsAt *time.Time, expiresAt *time.Time) error {
	err := t.taskAccessor.UpdateStartsAt(id, startsAt)
	if err != nil {
		return err
	}

	err = t.taskAccessor.UpdateExpiresAt(id, expiresAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTitle changes a task title.
func (t *Task) UpdateTitle(id uint, title string) error {
	err := t.taskAccessor.UpdateTitle(id, title)
	if err != nil {
		return err
	}

	return nil
}

// Lanize makes a task TypeLane.
func (t *Task) Lanize(id uint) error {
	err := t.taskAccessor.UpdateType(id, uint(TypeLane))
	if err != nil {
		return err
	}

	return nil
}

// MoveToRoot moves a task to root directory.
func (t *Task) MoveToRoot(taskID uint) error {
	err := t.taskAccessor.DeleteAncestorTaskRelations(taskID)
	if err != nil {
		return err
	}

	return nil
}

// MoveToChild moves childTask under parentTask.
func (t *Task) MoveToChild(parentTaskID uint, childTaskID uint) error {
	if err := t.taskAccessor.DeleteAncestorTaskRelations(childTaskID); err != nil {
		return err
	}

	if err := t.taskAccessor.CreateTaskRelationsBetweenTasks(parentTaskID, childTaskID); err != nil {
		return err
	}

	return nil
}

// String makes TaskType compliant with Stringer interface.
func (t TaskType) String() string {
	switch t {
	case TypeTask:
		return "task"
	case TypeLane:
		return "lane"
	default:
		return "undefined"
	}
}
