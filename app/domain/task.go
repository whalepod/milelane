package domain

import (
	"time"

	"github.com/whalepod/milelane/app/domain/repository"
)

type TaskType uint

const (
	TypeTask TaskType = iota * 10
	TypeLane
)

type TaskAccessor interface {
	ListTree() (*[]repository.TreeableTask, error)
	FindTreeByID(id uint) (*repository.TreeableTask, error)
	Create(title string) (*repository.Task, error)
	UpdateCompletedAt(id uint, completedAt time.Time) error
	UpdateType(id uint, taskType uint) error
	DeleteAncestorTaskRelations(taskID uint) error
	CreateTaskRelationsBetweenTasks(parentTaskID uint, childTaskID uint) error
}

// Struct for domain, not for gorm.
type Task struct {
	taskAccessor TaskAccessor
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	CompletedAt  string `json:"completed_at"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Depth        uint   `json:"depth"`
	Children     []Task `json:"children"`
}

func NewTask(ta TaskAccessor) (*Task, error) {
	var t Task
	t.taskAccessor = ta

	return &t, nil
}

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
			CreatedAt: treeableTask.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: treeableTask.UpdatedAt.Format("2006-01-02 15:04:05"),
			Depth:     treeableTask.Depth,
		}
		if treeableTask.CompletedAt != nil {
			task.CompletedAt = treeableTask.CompletedAt.Format("2006-01-02 15:04:05")
		}
		task = task.mapChildren(treeableTask.Children)
		tasks = append(tasks, task)
	}

	return &tasks, nil
}

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
		CreatedAt: treeableTask.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: treeableTask.UpdatedAt.Format("2006-01-02 15:04:05"),
		Depth:     treeableTask.Depth,
	}
	if treeableTask.CompletedAt != nil {
		task.CompletedAt = treeableTask.CompletedAt.Format("2006-01-02 15:04:05")
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
			CreatedAt: treeableTask.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: treeableTask.UpdatedAt.Format("2006-01-02 15:04:05"),
			Depth:     treeableTask.Depth,
		}
		if treeableTask.CompletedAt != nil {
			task.CompletedAt = treeableTask.CompletedAt.Format("2006-01-02 15:04:05")
		}
		task = task.mapChildren(treeableTask.Children)
		children = append(children, task)
	}
	(*t).Children = children

	return *t
}

func (t *Task) Create(title string) (*Task, error) {
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

	return &task, nil
}

func (t *Task) Complete(id uint) error {
	err := t.taskAccessor.UpdateCompletedAt(id, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) Lanize(id uint) error {
	err := t.taskAccessor.UpdateType(id, uint(TypeLane))
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) MoveToRoot(taskID uint) error {
	err := t.taskAccessor.DeleteAncestorTaskRelations(taskID)
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) MoveToChild(parentTaskID uint, childTaskID uint) error {
	if err := t.taskAccessor.DeleteAncestorTaskRelations(childTaskID); err != nil {
		return err
	}

	if err := t.taskAccessor.CreateTaskRelationsBetweenTasks(parentTaskID, childTaskID); err != nil {
		return err
	}

	return nil
}

// TaskType is compliant with Stringer interface.
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
