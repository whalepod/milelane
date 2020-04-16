package domain

import (
	"reflect"
	"testing"
	"time"

	"golang.org/x/xerrors"

	"github.com/whalepod/milelane/app/domain/repository"
)

type TaskAccessorMock struct{}

func (*TaskAccessorMock) ListTree() (*[]repository.TreeableTask, error) {
	tasks := []repository.TreeableTask{
		{
			Task: repository.Task{
				ID:          1,
				Title:       "trunk",
				Type:        uint(TypeLane),
				CompletedAt: &now,
				StartsAt:    &now,
				ExpiresAt:   &now,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			Depth: 1,
			Children: []repository.TreeableTask{
				{
					Task: repository.Task{
						ID:          2,
						Title:       "branch",
						Type:        uint(TypeTask),
						CompletedAt: &now,
						StartsAt:    &now,
						ExpiresAt:   &now,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					Depth: 2,
					Children: []repository.TreeableTask{
						{
							Task: repository.Task{
								ID:          3,
								Title:       "leaf",
								Type:        uint(TypeTask),
								CompletedAt: &now,
								StartsAt:    &now,
								ExpiresAt:   &now,
								CreatedAt:   now,
								UpdatedAt:   now,
							},
							Depth: 3,
						},
						{
							Task: repository.Task{
								ID:          5,
								Title:       "leaf-2",
								Type:        uint(TypeTask),
								CompletedAt: &now,
								StartsAt:    &now,
								ExpiresAt:   &now,
								CreatedAt:   now,
								UpdatedAt:   now,
							},
							Depth: 3,
						},
					},
				},
				{
					Task: repository.Task{
						ID:          4,
						Title:       "branch-2",
						Type:        uint(TypeTask),
						CompletedAt: &now,
						StartsAt:    &now,
						ExpiresAt:   &now,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					Depth: 2,
				},
			},
		},
	}
	return &tasks, nil
}

func (*TaskAccessorMock) ListTreeByDeviceUUID(deviceUUID string) (*[]repository.TreeableTask, error) {
	tasks := []repository.TreeableTask{
		{
			Task: repository.Task{
				ID:          1,
				Title:       "trunk",
				Type:        uint(TypeLane),
				CompletedAt: &now,
				StartsAt:    &now,
				ExpiresAt:   &now,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			Depth: 1,
			Children: []repository.TreeableTask{
				{
					Task: repository.Task{
						ID:          2,
						Title:       "branch",
						Type:        uint(TypeTask),
						CompletedAt: &now,
						StartsAt:    &now,
						ExpiresAt:   &now,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					Depth: 2,
					Children: []repository.TreeableTask{
						{
							Task: repository.Task{
								ID:          3,
								Title:       "leaf",
								Type:        uint(TypeTask),
								CompletedAt: &now,
								StartsAt:    &now,
								ExpiresAt:   &now,
								CreatedAt:   now,
								UpdatedAt:   now,
							},
							Depth: 3,
						},
						{
							Task: repository.Task{
								ID:          5,
								Title:       "leaf-2",
								Type:        uint(TypeTask),
								CompletedAt: &now,
								StartsAt:    &now,
								ExpiresAt:   &now,
								CreatedAt:   now,
								UpdatedAt:   now,
							},
							Depth: 3,
						},
					},
				},
				{
					Task: repository.Task{
						ID:          4,
						Title:       "branch-2",
						Type:        uint(TypeTask),
						CompletedAt: &now,
						StartsAt:    &now,
						ExpiresAt:   &now,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					Depth: 2,
				},
			},
		},
	}
	return &tasks, nil
}

func (*TaskAccessorMock) FindTreeByID(id uint) (*repository.TreeableTask, error) {
	task := repository.TreeableTask{
		Task: repository.Task{
			ID:          1,
			Title:       "trunk",
			Type:        uint(TypeLane),
			CompletedAt: &now,
			StartsAt:    &now,
			ExpiresAt:   &now,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		Depth: 1,
		Children: []repository.TreeableTask{
			{
				Task: repository.Task{
					ID:          2,
					Title:       "branch",
					Type:        uint(TypeTask),
					CompletedAt: &now,
					StartsAt:    &now,
					ExpiresAt:   &now,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
				Depth: 2,
				Children: []repository.TreeableTask{
					{
						Task: repository.Task{
							ID:          3,
							Title:       "leaf",
							Type:        uint(TypeTask),
							CompletedAt: &now,
							StartsAt:    &now,
							ExpiresAt:   &now,
							CreatedAt:   now,
							UpdatedAt:   now,
						},
						Depth: 3,
					},
					{
						Task: repository.Task{
							ID:          5,
							Title:       "leaf-2",
							Type:        uint(TypeTask),
							CompletedAt: &now,
							StartsAt:    &now,
							ExpiresAt:   &now,
							CreatedAt:   now,
							UpdatedAt:   now,
						},
						Depth: 3,
					},
				},
			},
			{
				Task: repository.Task{
					ID:          4,
					Title:       "branch-2",
					Type:        uint(TypeTask),
					CompletedAt: &now,
					StartsAt:    &now,
					ExpiresAt:   &now,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
				Depth: 2,
			},
		},
	}
	return &task, nil
}

func (*TaskAccessorMock) Create(title string) (*repository.Task, error) {
	var task repository.Task
	return &task, nil
}

func (*TaskAccessorMock) UpdateCompletedAt(id uint, completedAt time.Time) error {
	return nil
}

func (*TaskAccessorMock) UpdateType(id uint, taskType uint) error {
	return nil
}

func (*TaskAccessorMock) DeleteAncestorTaskRelations(taskID uint) error {
	return nil
}

func (*TaskAccessorMock) CreateTaskRelationsBetweenTasks(parentTaskID uint, childTaskID uint) error {
	return nil
}

func (*TaskAccessorMock) CreateDeviceTask(deviceUUID string, taskID uint) (*repository.DeviceTask, error) {
	var deviceTask repository.DeviceTask
	return &deviceTask, nil
}

type TaskAccessorErrorMock struct{}

func (*TaskAccessorErrorMock) ListTree() (*[]repository.TreeableTask, error) {
	return nil, xerrors.New("error mock called")
}

func (*TaskAccessorErrorMock) ListTreeByDeviceUUID(deviceUUID string) (*[]repository.TreeableTask, error) {
	return nil, xerrors.New("error mock called")
}

func (*TaskAccessorErrorMock) FindTreeByID(id uint) (*repository.TreeableTask, error) {
	return nil, xerrors.New("error mock called")
}

func (*TaskAccessorErrorMock) Create(title string) (*repository.Task, error) {
	return nil, xerrors.New("error mock called")
}

func (*TaskAccessorErrorMock) UpdateCompletedAt(id uint, completedAt time.Time) error {
	return xerrors.New("error mock called")
}

func (*TaskAccessorErrorMock) UpdateType(id uint, taskType uint) error {
	return xerrors.New("error mock called")
}

func (*TaskAccessorErrorMock) DeleteAncestorTaskRelations(taskID uint) error {
	return xerrors.New("error mock called")
}

func (*TaskAccessorErrorMock) CreateTaskRelationsBetweenTasks(parentTaskID uint, childTaskID uint) error {
	return xerrors.New("error mock called")
}

func (*TaskAccessorErrorMock) CreateDeviceTask(deviceUUID string, taskID uint) (*repository.DeviceTask, error) {
	return nil, xerrors.New("error mock called")
}

type TaskAccessorMoveToChildErrorMock struct{}

func (*TaskAccessorMoveToChildErrorMock) ListTree() (*[]repository.TreeableTask, error) {
	tasks := []repository.TreeableTask{
		{
			Task: repository.Task{
				ID:          1,
				Title:       "trunk",
				Type:        uint(TypeLane),
				CompletedAt: &now,
				StartsAt:    &now,
				ExpiresAt:   &now,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			Depth: 1,
			Children: []repository.TreeableTask{
				{
					Task: repository.Task{
						ID:          2,
						Title:       "branch",
						Type:        uint(TypeTask),
						CompletedAt: &now,
						StartsAt:    &now,
						ExpiresAt:   &now,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					Depth: 2,
					Children: []repository.TreeableTask{
						{
							Task: repository.Task{
								ID:          3,
								Title:       "leaf",
								Type:        uint(TypeTask),
								CompletedAt: &now,
								StartsAt:    &now,
								ExpiresAt:   &now,
								CreatedAt:   now,
								UpdatedAt:   now,
							},
							Depth: 3,
						},
						{
							Task: repository.Task{
								ID:          5,
								Title:       "leaf-2",
								Type:        uint(TypeTask),
								CompletedAt: &now,
								StartsAt:    &now,
								ExpiresAt:   &now,
								CreatedAt:   now,
								UpdatedAt:   now,
							},
							Depth: 3,
						},
					},
				},
				{
					Task: repository.Task{
						ID:          4,
						Title:       "branch-2",
						Type:        uint(TypeTask),
						CompletedAt: &now,
						StartsAt:    &now,
						ExpiresAt:   &now,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					Depth: 2,
				},
			},
		},
	}
	return &tasks, nil
}

func (*TaskAccessorMoveToChildErrorMock) ListTreeByDeviceUUID(deviceUUID string) (*[]repository.TreeableTask, error) {
	tasks := []repository.TreeableTask{
		{
			Task: repository.Task{
				ID:          1,
				Title:       "trunk",
				Type:        uint(TypeLane),
				CompletedAt: &now,
				StartsAt:    &now,
				ExpiresAt:   &now,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			Depth: 1,
			Children: []repository.TreeableTask{
				{
					Task: repository.Task{
						ID:          2,
						Title:       "branch",
						Type:        uint(TypeTask),
						CompletedAt: &now,
						StartsAt:    &now,
						ExpiresAt:   &now,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					Depth: 2,
					Children: []repository.TreeableTask{
						{
							Task: repository.Task{
								ID:          3,
								Title:       "leaf",
								Type:        uint(TypeTask),
								CompletedAt: &now,
								StartsAt:    &now,
								ExpiresAt:   &now,
								CreatedAt:   now,
								UpdatedAt:   now,
							},
							Depth: 3,
						},
						{
							Task: repository.Task{
								ID:          5,
								Title:       "leaf-2",
								Type:        uint(TypeTask),
								CompletedAt: &now,
								StartsAt:    &now,
								ExpiresAt:   &now,
								CreatedAt:   now,
								UpdatedAt:   now,
							},
							Depth: 3,
						},
					},
				},
				{
					Task: repository.Task{
						ID:          4,
						Title:       "branch-2",
						Type:        uint(TypeTask),
						CompletedAt: &now,
						StartsAt:    &now,
						ExpiresAt:   &now,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					Depth: 2,
				},
			},
		},
	}
	return &tasks, nil
}

func (*TaskAccessorMoveToChildErrorMock) FindTreeByID(id uint) (*repository.TreeableTask, error) {
	task := repository.TreeableTask{
		Task: repository.Task{
			ID:          1,
			Title:       "trunk",
			Type:        uint(TypeLane),
			CompletedAt: &now,
			StartsAt:    &now,
			ExpiresAt:   &now,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		Depth: 1,
		Children: []repository.TreeableTask{
			{
				Task: repository.Task{
					ID:          2,
					Title:       "branch",
					Type:        uint(TypeTask),
					CompletedAt: &now,
					StartsAt:    &now,
					ExpiresAt:   &now,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
				Depth: 2,
				Children: []repository.TreeableTask{
					{
						Task: repository.Task{
							ID:          3,
							Title:       "leaf",
							Type:        uint(TypeTask),
							CompletedAt: &now,
							StartsAt:    &now,
							ExpiresAt:   &now,
							CreatedAt:   now,
							UpdatedAt:   now,
						},
						Depth: 3,
					},
					{
						Task: repository.Task{
							ID:          5,
							Title:       "leaf-2",
							Type:        uint(TypeTask),
							CompletedAt: &now,
							StartsAt:    &now,
							ExpiresAt:   &now,
							CreatedAt:   now,
							UpdatedAt:   now,
						},
						Depth: 3,
					},
				},
			},
			{
				Task: repository.Task{
					ID:          4,
					Title:       "branch-2",
					Type:        uint(TypeTask),
					CompletedAt: &now,
					StartsAt:    &now,
					ExpiresAt:   &now,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
				Depth: 2,
			},
		},
	}
	return &task, nil
}

func (*TaskAccessorMoveToChildErrorMock) Create(title string) (*repository.Task, error) {
	var task repository.Task
	return &task, nil
}

func (*TaskAccessorMoveToChildErrorMock) UpdateCompletedAt(id uint, completedAt time.Time) error {
	return nil
}

func (*TaskAccessorMoveToChildErrorMock) UpdateType(id uint, taskType uint) error {
	return nil
}

func (*TaskAccessorMoveToChildErrorMock) DeleteAncestorTaskRelations(taskID uint) error {
	return nil
}

func (*TaskAccessorMoveToChildErrorMock) CreateTaskRelationsBetweenTasks(parentTaskID uint, childTaskID uint) error {
	return xerrors.New("error mock called")
}

func (*TaskAccessorMoveToChildErrorMock) CreateDeviceTask(deviceUUID string, taskID uint) (*repository.DeviceTask, error) {
	var deviceTask repository.DeviceTask
	return &deviceTask, nil
}

type TaskAccessorCreateDeviceTaskErrorMock struct{}

func (*TaskAccessorCreateDeviceTaskErrorMock) ListTree() (*[]repository.TreeableTask, error) {
	tasks := []repository.TreeableTask{}
	return &tasks, nil
}

func (*TaskAccessorCreateDeviceTaskErrorMock) ListTreeByDeviceUUID(deviceUUID string) (*[]repository.TreeableTask, error) {
	tasks := []repository.TreeableTask{}
	return &tasks, nil
}

func (*TaskAccessorCreateDeviceTaskErrorMock) FindTreeByID(id uint) (*repository.TreeableTask, error) {
	task := repository.TreeableTask{}
	return &task, nil
}

func (*TaskAccessorCreateDeviceTaskErrorMock) Create(title string) (*repository.Task, error) {
	var task repository.Task
	return &task, nil
}

func (*TaskAccessorCreateDeviceTaskErrorMock) UpdateCompletedAt(id uint, completedAt time.Time) error {
	return nil
}

func (*TaskAccessorCreateDeviceTaskErrorMock) UpdateType(id uint, taskType uint) error {
	return nil
}

func (*TaskAccessorCreateDeviceTaskErrorMock) DeleteAncestorTaskRelations(taskID uint) error {
	return nil
}

func (*TaskAccessorCreateDeviceTaskErrorMock) CreateTaskRelationsBetweenTasks(parentTaskID uint, childTaskID uint) error {
	return nil
}

func (*TaskAccessorCreateDeviceTaskErrorMock) CreateDeviceTask(deviceUUID string, taskID uint) (*repository.DeviceTask, error) {
	return nil, xerrors.New("error mock called")
}

func TestList(t *testing.T) {
	var taskAccessor TaskAccessorMock
	// Avoid conflict with comparing task object.
	ta, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	tasks, err := ta.List()
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	// Initialize task to be compared.
	expectedTasks := []Task{
		{
			ID:          1,
			Title:       "trunk",
			Type:        TypeLane.String(),
			CompletedAt: (&now).Format("2006-01-02 15:04:05"),
			StartsAt:    (&now).Format("2006-01-02 15:04:05"),
			ExpiresAt:   (&now).Format("2006-01-02 15:04:05"),
			CreatedAt:   now.Format("2006-01-02 15:04:05"),
			UpdatedAt:   now.Format("2006-01-02 15:04:05"),
			Depth:       1,
			Children: []Task{
				{
					ID:          2,
					Title:       "branch",
					Type:        TypeTask.String(),
					CompletedAt: (&now).Format("2006-01-02 15:04:05"),
					StartsAt:    (&now).Format("2006-01-02 15:04:05"),
					ExpiresAt:   (&now).Format("2006-01-02 15:04:05"),
					CreatedAt:   now.Format("2006-01-02 15:04:05"),
					UpdatedAt:   now.Format("2006-01-02 15:04:05"),
					Depth:       2,
					Children: []Task{
						{
							ID:          3,
							Title:       "leaf",
							Type:        TypeTask.String(),
							CompletedAt: (&now).Format("2006-01-02 15:04:05"),
							StartsAt:    (&now).Format("2006-01-02 15:04:05"),
							ExpiresAt:   (&now).Format("2006-01-02 15:04:05"),
							CreatedAt:   now.Format("2006-01-02 15:04:05"),
							UpdatedAt:   now.Format("2006-01-02 15:04:05"),
							Depth:       3,
						},
						{
							ID:          5,
							Title:       "leaf-2",
							Type:        TypeTask.String(),
							CompletedAt: (&now).Format("2006-01-02 15:04:05"),
							StartsAt:    (&now).Format("2006-01-02 15:04:05"),
							ExpiresAt:   (&now).Format("2006-01-02 15:04:05"),
							CreatedAt:   now.Format("2006-01-02 15:04:05"),
							UpdatedAt:   now.Format("2006-01-02 15:04:05"),
							Depth:       3,
						},
					},
				},
				{
					ID:          4,
					Title:       "branch-2",
					Type:        TypeTask.String(),
					CompletedAt: (&now).Format("2006-01-02 15:04:05"),
					StartsAt:    (&now).Format("2006-01-02 15:04:05"),
					ExpiresAt:   (&now).Format("2006-01-02 15:04:05"),
					CreatedAt:   now.Format("2006-01-02 15:04:05"),
					UpdatedAt:   now.Format("2006-01-02 15:04:05"),
					Depth:       2,
				},
			},
		},
	}

	if !reflect.DeepEqual(expectedTasks, *tasks) {
		t.Fatalf("Got wrong uncompleted task. Want: %v, Got: %v,  ", expectedTasks, *tasks)
	}

	t.Log("Success.")
}

func TestListError(t *testing.T) {
	var taskAccessor TaskAccessorErrorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	_, err = task.List()
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success: Got expected err.")
}

func TestFind(t *testing.T) {
	var taskAccessor TaskAccessorMock
	// Avoid conflict with comparing task object.
	ta, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	var id uint = 1
	task, err := ta.Find(id)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	// Initialize task to be compared.
	expectedTask := Task{
		ID:          1,
		Title:       "trunk",
		Type:        TypeLane.String(),
		CompletedAt: (&now).Format("2006-01-02 15:04:05"),
		StartsAt:    (&now).Format("2006-01-02 15:04:05"),
		ExpiresAt:   (&now).Format("2006-01-02 15:04:05"),
		CreatedAt:   now.Format("2006-01-02 15:04:05"),
		UpdatedAt:   now.Format("2006-01-02 15:04:05"),
		Depth:       1,
		Children: []Task{
			{
				ID:          2,
				Title:       "branch",
				Type:        TypeTask.String(),
				CompletedAt: (&now).Format("2006-01-02 15:04:05"),
				StartsAt:    (&now).Format("2006-01-02 15:04:05"),
				ExpiresAt:   (&now).Format("2006-01-02 15:04:05"),
				CreatedAt:   now.Format("2006-01-02 15:04:05"),
				UpdatedAt:   now.Format("2006-01-02 15:04:05"),
				Depth:       2,
				Children: []Task{
					{
						ID:          3,
						Title:       "leaf",
						Type:        TypeTask.String(),
						CompletedAt: (&now).Format("2006-01-02 15:04:05"),
						StartsAt:    (&now).Format("2006-01-02 15:04:05"),
						ExpiresAt:   (&now).Format("2006-01-02 15:04:05"),
						CreatedAt:   now.Format("2006-01-02 15:04:05"),
						UpdatedAt:   now.Format("2006-01-02 15:04:05"),
						Depth:       3,
					},
					{
						ID:          5,
						Title:       "leaf-2",
						Type:        TypeTask.String(),
						CompletedAt: (&now).Format("2006-01-02 15:04:05"),
						StartsAt:    (&now).Format("2006-01-02 15:04:05"),
						ExpiresAt:   (&now).Format("2006-01-02 15:04:05"),
						CreatedAt:   now.Format("2006-01-02 15:04:05"),
						UpdatedAt:   now.Format("2006-01-02 15:04:05"),
						Depth:       3,
					},
				},
			},
			{
				ID:          4,
				Title:       "branch-2",
				Type:        TypeTask.String(),
				CompletedAt: (&now).Format("2006-01-02 15:04:05"),
				StartsAt:    (&now).Format("2006-01-02 15:04:05"),
				ExpiresAt:   (&now).Format("2006-01-02 15:04:05"),
				CreatedAt:   now.Format("2006-01-02 15:04:05"),
				UpdatedAt:   now.Format("2006-01-02 15:04:05"),
				Depth:       2,
			},
		},
	}

	if !reflect.DeepEqual(expectedTask, *task) {
		t.Fatalf("Got wrong uncompleted task. Want: %v, Got: %v,  ", expectedTask, *task)
	}

	t.Log("Success.")
}

func TestFindError(t *testing.T) {
	var taskAccessor TaskAccessorErrorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	var id uint = 1
	_, err = task.Find(id)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success: Got expected err.")
}

func TestCreate(t *testing.T) {
	var taskAccessor TaskAccessorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	title := "Test input."
	_, err = task.Create(title)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestCreateError(t *testing.T) {
	var taskAccessor TaskAccessorErrorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	title := "Test wrong input."
	_, err = task.Create(title)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success: Got expected err.")
}

func TestComplete(t *testing.T) {
	var taskAccessor TaskAccessorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.Complete(1)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestCompleteError(t *testing.T) {
	var taskAccessor TaskAccessorErrorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.Complete(1)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success: Got expected err.")
}

func TestLanize(t *testing.T) {
	var taskAccessor TaskAccessorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.Lanize(1)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestLanizeError(t *testing.T) {
	var taskAccessor TaskAccessorErrorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.Lanize(1)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success: Got expected err.")
}

func TestMoveToRoot(t *testing.T) {
	var taskAccessor TaskAccessorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.MoveToRoot(1)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestMoveToRootError(t *testing.T) {
	var taskAccessor TaskAccessorErrorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.MoveToRoot(1)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success: Got expected err.")
}

func TestMoveToChild(t *testing.T) {
	var taskAccessor TaskAccessorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.MoveToChild(1, 2)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestMoveToChildErrorOnDeleteAncestorTaskRelations(t *testing.T) {
	var taskAccessor TaskAccessorErrorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.MoveToChild(1, 2)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success.")
}

func TestMoveToChildErrorOnCreateTaskRelationsBetweenTasks(t *testing.T) {
	var taskAccessor TaskAccessorMoveToChildErrorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.MoveToChild(1, 2)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success.")
}

func TestTaskTypeString(t *testing.T) {
	taskType := TypeTask
	if taskType.String() != "task" {
		t.Fatalf("TypeTask conversion to string failed, got response: %s", taskType.String())
	}

	taskType = TypeLane
	if taskType.String() != "lane" {
		t.Fatalf("TypeTask conversion to string failed, got response: %s", taskType.String())
	}

	taskType = 999
	if taskType.String() != "undefined" {
		t.Fatalf("TypeTask conversion to string failed, got response: %s", taskType.String())
	}

	t.Log("Success.")
}
