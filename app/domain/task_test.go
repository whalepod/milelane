package domain

import (
	"reflect"
	"testing"
	"time"

	"golang.org/x/xerrors"

	"github.com/whalepod/milelane/app/domain/repository"
)

type TaskAccessorMock struct {
	ErrorTarget string
}

func (mock *TaskAccessorMock) ListTree() (*[]repository.TreeableTask, error) {
	if mock.ErrorTarget == "ListTree" {
		return nil, xerrors.New("error mock called")
	}

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

func (mock *TaskAccessorMock) ListTreeByDeviceUUID(deviceUUID string) (*[]repository.TreeableTask, error) {
	if mock.ErrorTarget == "ListTreeByDeviceUUID" {
		return nil, xerrors.New("error mock called")
	}

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

func (mock *TaskAccessorMock) FindTreeByID(id uint) (*repository.TreeableTask, error) {
	if mock.ErrorTarget == "FindTreeByID" {
		return nil, xerrors.New("error mock called")
	}

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

func (mock *TaskAccessorMock) Create(title string) (*repository.Task, error) {
	if mock.ErrorTarget == "Create" {
		return nil, xerrors.New("error mock called")
	}

	var task repository.Task
	return &task, nil
}

func (mock *TaskAccessorMock) UpdateCompletedAt(id uint, completedAt time.Time) error {
	if mock.ErrorTarget == "UpdateCompletedAt" {
		return xerrors.New("error mock called")
	}

	return nil
}

func (mock *TaskAccessorMock) UpdateStartsAt(id uint, startsAt *time.Time) error {
	if mock.ErrorTarget == "UpdateStartsAt" {
		return xerrors.New("error mock called")
	}

	return nil
}

func (mock *TaskAccessorMock) UpdateExpiresAt(id uint, expiresAt *time.Time) error {
	if mock.ErrorTarget == "UpdateExpiresAt" {
		return xerrors.New("error mock called")
	}

	return nil
}

func (mock *TaskAccessorMock) UpdateTitle(id uint, title string) error {
	if mock.ErrorTarget == "UpdateTitle" {
		return xerrors.New("error mock called")
	}

	return nil
}

func (mock *TaskAccessorMock) UpdateType(id uint, taskType uint) error {
	if mock.ErrorTarget == "UpdateType" {
		return xerrors.New("error mock called")
	}

	return nil
}

func (mock *TaskAccessorMock) DeleteAncestorTaskRelations(taskID uint) error {
	if mock.ErrorTarget == "DeleteAncestorTaskRelations" {
		return xerrors.New("error mock called")
	}

	return nil
}

func (mock *TaskAccessorMock) CreateTaskRelationsBetweenTasks(parentTaskID uint, childTaskID uint) error {
	if mock.ErrorTarget == "CreateTaskRelationsBetweenTasks" {
		return xerrors.New("error mock called")
	}

	return nil
}

func (mock *TaskAccessorMock) CreateDeviceTask(deviceUUID string, taskID uint) (*repository.DeviceTask, error) {
	if mock.ErrorTarget == "CreateDeviceTask" {
		return nil, xerrors.New("error mock called")
	}

	var deviceTask repository.DeviceTask
	return &deviceTask, nil
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
			CompletedAt: (&now).In(time.Local).Format("2006-01-02 15:04:05"),
			StartsAt:    (&now).In(time.Local).Format("2006-01-02 15:04:05"),
			ExpiresAt:   (&now).In(time.Local).Format("2006-01-02 15:04:05"),
			CreatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
			UpdatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
			Depth:       1,
			Children: []Task{
				{
					ID:          2,
					Title:       "branch",
					Type:        TypeTask.String(),
					CompletedAt: (&now).In(time.Local).Format("2006-01-02 15:04:05"),
					StartsAt:    (&now).In(time.Local).Format("2006-01-02 15:04:05"),
					ExpiresAt:   (&now).In(time.Local).Format("2006-01-02 15:04:05"),
					CreatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
					UpdatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
					Depth:       2,
					Children: []Task{
						{
							ID:          3,
							Title:       "leaf",
							Type:        TypeTask.String(),
							CompletedAt: (&now).In(time.Local).Format("2006-01-02 15:04:05"),
							StartsAt:    (&now).In(time.Local).Format("2006-01-02 15:04:05"),
							ExpiresAt:   (&now).In(time.Local).Format("2006-01-02 15:04:05"),
							CreatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
							UpdatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
							Depth:       3,
						},
						{
							ID:          5,
							Title:       "leaf-2",
							Type:        TypeTask.String(),
							CompletedAt: (&now).In(time.Local).Format("2006-01-02 15:04:05"),
							StartsAt:    (&now).In(time.Local).Format("2006-01-02 15:04:05"),
							ExpiresAt:   (&now).In(time.Local).Format("2006-01-02 15:04:05"),
							CreatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
							UpdatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
							Depth:       3,
						},
					},
				},
				{
					ID:          4,
					Title:       "branch-2",
					Type:        TypeTask.String(),
					CompletedAt: (&now).In(time.Local).Format("2006-01-02 15:04:05"),
					StartsAt:    (&now).In(time.Local).Format("2006-01-02 15:04:05"),
					ExpiresAt:   (&now).In(time.Local).Format("2006-01-02 15:04:05"),
					CreatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
					UpdatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
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
	taskAccessor := TaskAccessorMock{ErrorTarget: "ListTree"}
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
		CompletedAt: (&now).In(time.Local).Format("2006-01-02 15:04:05"),
		StartsAt:    (&now).In(time.Local).Format("2006-01-02 15:04:05"),
		ExpiresAt:   (&now).In(time.Local).Format("2006-01-02 15:04:05"),
		CreatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
		UpdatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
		Depth:       1,
		Children: []Task{
			{
				ID:          2,
				Title:       "branch",
				Type:        TypeTask.String(),
				CompletedAt: (&now).In(time.Local).Format("2006-01-02 15:04:05"),
				StartsAt:    (&now).In(time.Local).Format("2006-01-02 15:04:05"),
				ExpiresAt:   (&now).In(time.Local).Format("2006-01-02 15:04:05"),
				CreatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
				UpdatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
				Depth:       2,
				Children: []Task{
					{
						ID:          3,
						Title:       "leaf",
						Type:        TypeTask.String(),
						CompletedAt: (&now).In(time.Local).Format("2006-01-02 15:04:05"),
						StartsAt:    (&now).In(time.Local).Format("2006-01-02 15:04:05"),
						ExpiresAt:   (&now).In(time.Local).Format("2006-01-02 15:04:05"),
						CreatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
						UpdatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
						Depth:       3,
					},
					{
						ID:          5,
						Title:       "leaf-2",
						Type:        TypeTask.String(),
						CompletedAt: (&now).In(time.Local).Format("2006-01-02 15:04:05"),
						StartsAt:    (&now).In(time.Local).Format("2006-01-02 15:04:05"),
						ExpiresAt:   (&now).In(time.Local).Format("2006-01-02 15:04:05"),
						CreatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
						UpdatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
						Depth:       3,
					},
				},
			},
			{
				ID:          4,
				Title:       "branch-2",
				Type:        TypeTask.String(),
				CompletedAt: (&now).In(time.Local).Format("2006-01-02 15:04:05"),
				StartsAt:    (&now).In(time.Local).Format("2006-01-02 15:04:05"),
				ExpiresAt:   (&now).In(time.Local).Format("2006-01-02 15:04:05"),
				CreatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
				UpdatedAt:   now.In(time.Local).Format("2006-01-02 15:04:05"),
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
	taskAccessor := TaskAccessorMock{ErrorTarget: "FindTreeByID"}
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
	taskAccessor := TaskAccessorMock{ErrorTarget: "Create"}
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
	taskAccessor := TaskAccessorMock{ErrorTarget: "UpdateCompletedAt"}
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

func TestUpdateTerm(t *testing.T) {
	var taskAccessor TaskAccessorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.UpdateTerm(1, &now, &now)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestUpdateTermErrorOnUpdateStartsAt(t *testing.T) {
	taskAccessor := TaskAccessorMock{ErrorTarget: "UpdateStartsAt"}
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.UpdateTerm(1, &now, &now)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success: Got expected err.")
}

func TestUpdateTermErrorOnUpdateExpiresAt(t *testing.T) {
	taskAccessor := TaskAccessorMock{ErrorTarget: "UpdateExpiresAt"}
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.UpdateTerm(1, &now, &now)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success: Got expected err.")
}

func TestUpdateTitle(t *testing.T) {
	var taskAccessor TaskAccessorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.UpdateTitle(1, "Update test")
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestUpdateTitleError(t *testing.T) {
	taskAccessor := TaskAccessorMock{ErrorTarget: "UpdateTitle"}
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.UpdateTitle(1, "Update test")
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
	taskAccessor := TaskAccessorMock{ErrorTarget: "UpdateType"}
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

func TestDelanize(t *testing.T) {
	var taskAccessor TaskAccessorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.Delanize(1)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestDelanizeError(t *testing.T) {
	taskAccessor := TaskAccessorMock{ErrorTarget: "UpdateType"}
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	err = task.Delanize(1)
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
	taskAccessor := TaskAccessorMock{ErrorTarget: "DeleteAncestorTaskRelations"}
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
	taskAccessor := TaskAccessorMock{ErrorTarget: "DeleteAncestorTaskRelations"}
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
	taskAccessor := TaskAccessorMock{ErrorTarget: "CreateTaskRelationsBetweenTasks"}
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
