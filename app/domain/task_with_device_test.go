package domain

import (
	"reflect"
	"testing"
)

func TestListByDeviceUUID(t *testing.T) {
	var taskAccessor TaskAccessorMock
	// Avoid conflict with comparing task object.
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	deviceUUID := "60982a48-9328-441f-805b-d3ab0cad9e1f"
	tasks, err := task.ListByDeviceUUID(deviceUUID)
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

func TestListByDeviceUUIDError(t *testing.T) {
	taskAccessor := TaskAccessorMock{ErrorTarget: "ListTreeByDeviceUUID"}
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	deviceUUID := "60982a48-9328-441f-805b-d3ab0cad9e1f"
	_, err = task.ListByDeviceUUID(deviceUUID)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success: Got expected err.")
}

func TestCreateWithDevice(t *testing.T) {
	var taskAccessor TaskAccessorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	deviceUUID := "60982a48-9328-441f-805b-d3ab0cad9e1f"
	title := "Test input."
	_, err = task.CreateWithDevice(deviceUUID, title)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestCreateDeviceTaskError(t *testing.T) {
	taskAccessor := TaskAccessorMock{ErrorTarget: "Create"}
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	deviceUUID := "60982a48-9328-441f-805b-d3ab0cad9e1f"
	title := "Test wrong input."
	_, err = task.CreateWithDevice(deviceUUID, title)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success: Got expected err.")
}

func TestCreateDeviceTaskErrorOnCreateDeviceTask(t *testing.T) {
	taskAccessor := TaskAccessorMock{ErrorTarget: "CreateDeviceTask"}
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	deviceUUID := "Wrong input."
	title := "Test input."
	_, err = task.CreateWithDevice(deviceUUID, title)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success: Got expected err.")
}
