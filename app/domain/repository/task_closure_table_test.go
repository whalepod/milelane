package repository

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	QueryTaskTreeSelect               = `SELECT tasks.id, tasks.title, tasks.type, tasks.completed_at, tasks.created_at, tasks.updated_at, max(descendant_relations.path_length) AS depth FROM tasks LEFT JOIN task_relations AS descendant_relations ON tasks.id = descendant_relations.descendant_id GROUP BY tasks.id, tasks.title, tasks.type, tasks.completed_at, tasks.created_at, tasks.updated_at, descendant_relations.descendant_id ORDER BY group_concat(descendant_relations.ancestor_id ORDER BY descendant_relations.path_length DESC), tasks.id`
	QueryTaskTreeSelectByDeviceUUID   = `SELECT tasks.id, tasks.title, tasks.type, tasks.completed_at, tasks.created_at, tasks.updated_at, max(descendant_relations.path_length) AS depth FROM tasks LEFT JOIN task_relations AS descendant_relations ON tasks.id = descendant_relations.descendant_id LEFT JOIN device_tasks ON tasks.id = device_tasks.task_id WHERE device_tasks.device_uuid = ? GROUP BY tasks.id, tasks.title, tasks.type, tasks.completed_at, tasks.created_at, tasks.updated_at, descendant_relations.descendant_id ORDER BY group_concat(descendant_relations.ancestor_id ORDER BY descendant_relations.path_length DESC), tasks.id`
	QueryTaskTreeSelectByTaskID       = `SELECT tasks.id, tasks.title, tasks.type, tasks.completed_at, tasks.created_at, tasks.updated_at, max(descendant_relations.path_length) AS depth FROM tasks LEFT JOIN task_relations AS descendant_relations ON tasks.id = descendant_relations.descendant_id WHERE ( tasks.id IN ( SELECT tasks.id FROM tasks LEFT JOIN task_relations ON tasks.id = task_relations.descendant_id WHERE ( task_relations.ancestor_id = ? ) ) ) GROUP BY tasks.id, tasks.title, tasks.type, tasks.completed_at, tasks.created_at, tasks.updated_at ORDER BY group_concat(descendant_relations.ancestor_id ORDER BY descendant_relations.path_length DESC), tasks.id`
	QueryTaskSelectSelfAndDescendants = `SELECT tasks.id, tasks.title, tasks.type, tasks.completed_at, tasks.created_at, tasks.updated_at FROM tasks LEFT JOIN task_relations AS descendant_relations ON tasks.id = descendant_relations.descendant_id WHERE descendant_relations.ancestor_id = ? GROUP BY tasks.id, tasks.title, tasks.type, tasks.completed_at, tasks.created_at, tasks.updated_at, descendant_relations.path_length ORDER BY descendant_relations.path_length asc`
	QueryTaskSelectSelfAndAncestors   = `SELECT tasks.id, tasks.title, tasks.type, tasks.completed_at, tasks.created_at, tasks.updated_at FROM tasks LEFT JOIN task_relations AS ancestor_relations ON tasks.id = ancestor_relations.ancestor_id WHERE ancestor_relations.descendant_id = ? GROUP BY tasks.id, tasks.title, tasks.type, tasks.completed_at, tasks.created_at, tasks.updated_at, ancestor_relations.path_length ORDER BY ancestor_relations.path_length asc`
	QueryTaskSelectPathLength         = `SELECT max(path_length) AS path_length FROM task_relations WHERE descendant_id = ?`
	QueryTaskRelationDeleteAncestors  = `DELETE FROM task_relations WHERE descendant_id IN (?,?,?) AND ancestor_id NOT IN (?,?,?)`
	QueryTaskSelectID2                = `SELECT * FROM "tasks" WHERE ("tasks"."id" = 2)`
	QueryTaskSelectID3                = `SELECT * FROM "tasks" WHERE ("tasks"."id" = 3)`
	QueryTaskRelationInsertRegex      = `INSERT INTO task_relations`
)

func TestListTree(t *testing.T) {
	expectedTreeableTasks := []TreeableTask{
		{
			Task: Task{
				ID:          1,
				Title:       "trunk",
				CompletedAt: &now,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			Depth: 1,
			Children: []TreeableTask{
				{
					Task: Task{
						ID:          2,
						Title:       "branch",
						CompletedAt: &now,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					Depth: 2,
					Children: []TreeableTask{
						{
							Task: Task{
								ID:          3,
								Title:       "leaf",
								CompletedAt: &now,
								CreatedAt:   now,
								UpdatedAt:   now,
							},
							Depth: 3,
						},
						{
							Task: Task{
								ID:          5,
								Title:       "leaf-2",
								CompletedAt: &now,
								CreatedAt:   now,
								UpdatedAt:   now,
							},
							Depth: 3,
						},
					},
				},
				{
					Task: Task{
						ID:          4,
						Title:       "branch-2",
						CompletedAt: &now,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					Depth: 2,
				},
			},
		},
	}

	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskTreeSelect)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at", "depth"}).
				AddRow("1", "trunk", now, now, now, 1).
				AddRow("2", "branch", now, now, now, 2).
				AddRow("3", "leaf", now, now, now, 3).
				AddRow("5", "leaf-2", now, now, now, 3).
				AddRow("4", "branch-2", now, now, now, 2))

	taskRepository := NewTask(db)
	tasks, err := taskRepository.ListTree()
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	if len(*tasks) == 0 {
		t.Fatalf("No treeable task mapped.")
	}

	if !reflect.DeepEqual(*tasks, expectedTreeableTasks) {
		t.Fatalf("Got wrong tasks. Want: %v, Got: %v", expectedTreeableTasks[0].Children, (*tasks)[0].Children)
	}

	t.Log("Success.")
}

func TestListTreeFailsScan(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskTreeSelect)).
		WillReturnError(fmt.Errorf("Task scanning failed"))

	taskRepository := NewTask(db)
	_, err := taskRepository.ListTree()
	if err.Error() != "Task scanning failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task scanning failed")
	}

	t.Log("Success.")
}

func TestListTreeByDeviceUUID(t *testing.T) {
	expectedTreeableTasks := []TreeableTask{
		{
			Task: Task{
				ID:          1,
				Title:       "trunk",
				CompletedAt: &now,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			Depth: 1,
			Children: []TreeableTask{
				{
					Task: Task{
						ID:          2,
						Title:       "branch",
						CompletedAt: &now,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					Depth: 2,
					Children: []TreeableTask{
						{
							Task: Task{
								ID:          3,
								Title:       "leaf",
								CompletedAt: &now,
								CreatedAt:   now,
								UpdatedAt:   now,
							},
							Depth: 3,
						},
						{
							Task: Task{
								ID:          5,
								Title:       "leaf-2",
								CompletedAt: &now,
								CreatedAt:   now,
								UpdatedAt:   now,
							},
							Depth: 3,
						},
					},
				},
				{
					Task: Task{
						ID:          4,
						Title:       "branch-2",
						CompletedAt: &now,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					Depth: 2,
				},
			},
		},
	}

	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskTreeSelectByDeviceUUID)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at", "depth"}).
				AddRow("1", "trunk", now, now, now, 1).
				AddRow("2", "branch", now, now, now, 2).
				AddRow("3", "leaf", now, now, now, 3).
				AddRow("5", "leaf-2", now, now, now, 3).
				AddRow("4", "branch-2", now, now, now, 2))

	taskRepository := NewTask(db)
	deviceUUID := "60982a48-9328-441f-805b-d3ab0cad9e1f"
	tasks, err := taskRepository.ListTreeByDeviceUUID(deviceUUID)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	if len(*tasks) == 0 {
		t.Fatalf("No treeable task mapped.")
	}

	if !reflect.DeepEqual(*tasks, expectedTreeableTasks) {
		t.Fatalf("Got wrong tasks. Want: %v, Got: %v", expectedTreeableTasks[0].Children, (*tasks)[0].Children)
	}

	t.Log("Success.")
}

func TestListTreeByDeviceUUIDFailsScan(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskTreeSelectByDeviceUUID)).
		WillReturnError(fmt.Errorf("Task scanning failed"))

	deviceUUID := "60982a48-9328-441f-805b-d3abnoresult"
	taskRepository := NewTask(db)
	_, err := taskRepository.ListTreeByDeviceUUID(deviceUUID)
	if err.Error() != "Task scanning failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task scanning failed")
	}

	t.Log("Success.")
}

func TestFindTreeByID(t *testing.T) {
	expectedTreeableTask := TreeableTask{
		Task: Task{
			ID:          2,
			Title:       "branch",
			CompletedAt: &now,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		Depth: 2,
		Children: []TreeableTask{
			{
				Task: Task{
					ID:          3,
					Title:       "leaf",
					CompletedAt: &now,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
				Depth: 3,
			},
			{
				Task: Task{
					ID:          5,
					Title:       "leaf-2",
					CompletedAt: &now,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
				Depth: 3,
			},
		},
	}

	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskTreeSelectByTaskID)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at", "depth"}).
				AddRow("2", "branch", now, now, now, 2).
				AddRow("3", "leaf", now, now, now, 3).
				AddRow("5", "leaf-2", now, now, now, 3))

	taskRepository := NewTask(db)
	var id uint = 2
	task, err := taskRepository.FindTreeByID(id)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	if !reflect.DeepEqual(*task, expectedTreeableTask) {
		t.Fatalf("Got wrong tasks. Want: %v, Got: %v", expectedTreeableTask, *task)
	}

	t.Log("Success.")
}

func TestFindTreeByIDFailsScan(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskTreeSelectByTaskID)).
		WillReturnError(fmt.Errorf("Task scanning failed"))

	var id uint = 2
	taskRepository := NewTask(db)
	_, err := taskRepository.FindTreeByID(id)
	if err.Error() != "Task scanning failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task scanning failed")
	}

	t.Log("Success.")
}

func TestFindTreeByIDErrorRecordNotFound(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskTreeSelectByTaskID)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at", "depth"}))

	var id uint = 9999
	taskRepository := NewTask(db)
	_, err := taskRepository.FindTreeByID(id)
	if err.Error() != "record not found" {
		t.Fatalf("Got %v\nwant %v", err, "record not found")
	}

	t.Log("Success.")
}

func TestListSelfAndDescendants(t *testing.T) {
	expectedTasks := []Task{
		{
			ID:          1,
			Title:       "trunk",
			CompletedAt: &now,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          2,
			Title:       "branch",
			CompletedAt: &now,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          3,
			Title:       "leaf",
			CompletedAt: &now,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndDescendants)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("1", "trunk", now, now, now).
				AddRow("2", "branch", now, now, now).
				AddRow("3", "leaf", now, now, now))

	taskRepository := NewTask(db)
	tasks, err := taskRepository.ListSelfAndDescendants(3)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	if len(*tasks) == 0 {
		t.Fatalf("No task mapped.")
	}

	if !reflect.DeepEqual(*tasks, expectedTasks) {
		t.Fatalf("Got wrong uncompleted task. Want: %v, Got: %v", *tasks, expectedTasks)
	}

	t.Log("Success.")
}

func TestListSelfAndDescendantsRecordNotFound(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndDescendants)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}))

	taskRepository := NewTask(db)
	_, err := taskRepository.ListSelfAndDescendants(3)
	if err.Error() != "record not found" {
		t.Fatalf("Got %v\nwant %v", err, "record not found")
	}

	t.Log("Success.")
}

func TestListSelfAndDescendantsCallsError(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndDescendants)).
		WillReturnError(fmt.Errorf("Task selection failed"))

	taskRepository := NewTask(db)
	_, err := taskRepository.ListSelfAndDescendants(3)
	if err.Error() != "Task selection failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task selection failed")
	}

	t.Log("Success.")
}

func TestListSelfAndAncestors(t *testing.T) {
	expectedTasks := []Task{
		{
			ID:          3,
			Title:       "leaf",
			CompletedAt: &now,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          2,
			Title:       "branch",
			CompletedAt: &now,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          1,
			Title:       "trunk",
			CompletedAt: &now,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndAncestors)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("3", "leaf", now, now, now).
				AddRow("2", "branch", now, now, now).
				AddRow("1", "trunk", now, now, now))

	taskRepository := NewTask(db)
	tasks, err := taskRepository.ListSelfAndAncestors(3)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	if len(*tasks) == 0 {
		t.Fatalf("No task mapped.")
	}

	if !reflect.DeepEqual(*tasks, expectedTasks) {
		t.Fatalf("Got wrong uncompleted task. Want: %v, Got: %v", *tasks, expectedTasks)
	}

	t.Log("Success.")
}

func TestListSelfAndAncestorsRecordNotFound(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndAncestors)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}))

	taskRepository := NewTask(db)
	_, err := taskRepository.ListSelfAndAncestors(3)
	if err.Error() != "record not found" {
		t.Fatalf("Got %v\nwant %v", err, "record not found")
	}

	t.Log("Success.")
}

func TestListSelfAndAncestorsCallsError(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndAncestors)).
		WillReturnError(fmt.Errorf("Task selection failed"))

	taskRepository := NewTask(db)
	_, err := taskRepository.ListSelfAndAncestors(3)
	if err.Error() != "Task selection failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task selection failed")
	}

	t.Log("Success.")
}

func TestGetLevel(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("3"))

	taskRepository := NewTask(db)
	level, err := taskRepository.GetLevel(3)
	if err != nil {
		t.Fatalf("Level returns unexpected error. Got: %v", err.Error())
	}

	if level != 3 {
		t.Fatalf("Got wrong result. Want: %v, Got: %v", level, 3)
	}

	t.Log("Success.")
}

func TestGetLevelError(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnError(fmt.Errorf("level acquisition failed"))

	taskRepository := NewTask(db)
	_, err := taskRepository.GetLevel(3)
	if err.Error() != "level acquisition failed" {
		t.Fatalf("Got %v\nwant %v", err, "level acquisition failed")
	}

	t.Log("Success.")
}

func TestDeleteAncestorTaskRelations(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndDescendants)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("1", "trunk", now, now, now).
				AddRow("2", "branch", now, now, now).
				AddRow("3", "leaf", now, now, now))

	mock.ExpectExec(regexp.QuoteMeta(QueryTaskRelationDeleteAncestors)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	taskRepository := NewTask(db)
	err := taskRepository.DeleteAncestorTaskRelations(3)
	if err != nil {
		t.Fatalf("Level returns unexpected error. Got: %v", err.Error())
	}

	t.Log("Success.")
}

func TestDeleteAncestorTaskRelationsFailsOnGetList(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndDescendants)).
		WillReturnError(fmt.Errorf("Task selection failed"))

	taskRepository := NewTask(db)
	err := taskRepository.DeleteAncestorTaskRelations(3)
	if err.Error() != "Task selection failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task selection failed")
	}

	t.Log("Success.")
}

func TestDeleteAncestorTaskRelationsFailOnDeletion(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndDescendants)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("1", "trunk", now, now, now).
				AddRow("2", "branch", now, now, now).
				AddRow("3", "leaf", now, now, now))

	mock.ExpectExec(regexp.QuoteMeta(QueryTaskRelationDeleteAncestors)).
		WillReturnError(fmt.Errorf("Task deletion failed"))

	taskRepository := NewTask(db)
	err := taskRepository.DeleteAncestorTaskRelations(3)
	if err.Error() != "Task deletion failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task deletion failed")
	}

	t.Log("Success.")
}

func TestCreateTaskRelationsBetweenTasks(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID2)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID3)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("3", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("2"))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndAncestors)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "trunk", now, now, now).
				AddRow("1", "root", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndDescendants)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("3", "branch", now, now, now).
				AddRow("4", "leaf", now, now, now))

	// First ancestor loop.
	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("1"))

	// First descendant loop.
	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("1"))

	// Second descendant loop.
	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("2"))

	// Second ancestor loop.
	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("2"))

	// Third descendant loop.
	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("1"))

	// Fourth descendant loop.
	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("2"))

	// Create section.
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO task_relations`)).
		WillReturnResult(sqlmock.NewResult(10, 4))

	taskRepository := NewTask(db)
	err := taskRepository.CreateTaskRelationsBetweenTasks(2, 3)
	if err != nil {
		t.Fatalf("Create TaskRelations returns unexpected error. Got: %v", err.Error())
	}

	t.Log("Success.")
}

func TestCreateTaskRelationsBetweenTasksFailOnFindParent(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID2)).
		WillReturnError(fmt.Errorf("Task selection failed"))

	taskRepository := NewTask(db)
	err := taskRepository.CreateTaskRelationsBetweenTasks(2, 3)
	if err.Error() != "Task selection failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task selection failed")
	}

	t.Log("Success.")
}

func TestCreateTaskRelationsBetweenTasksFailOnFindChild(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID2)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID3)).
		WillReturnError(fmt.Errorf("Task selection failed"))

	taskRepository := NewTask(db)
	err := taskRepository.CreateTaskRelationsBetweenTasks(2, 3)
	if err.Error() != "Task selection failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task selection failed")
	}

	t.Log("Success.")
}

func TestCreateTaskRelationsBetweenTasksFailOnGetParentLevel(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID2)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID3)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("3", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnError(fmt.Errorf("level acquisition failed"))

	taskRepository := NewTask(db)
	err := taskRepository.CreateTaskRelationsBetweenTasks(2, 3)
	if err.Error() != "level acquisition failed" {
		t.Fatalf("Got %v\nwant %v", err, "level acquisition failed")
	}

	t.Log("Success.")
}

func TestCreateTaskRelationsBetweenTasksFailOnGetAncestorList(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID2)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID3)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("3", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("2"))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndAncestors)).
		WillReturnError(fmt.Errorf("Task selection failed"))

	taskRepository := NewTask(db)
	err := taskRepository.CreateTaskRelationsBetweenTasks(2, 3)
	if err.Error() != "Task selection failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task selection failed")
	}

	t.Log("Success.")
}

func TestCreateTaskRelationsBetweenTasksFailOnGetDescendantList(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID2)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID3)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("3", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("2"))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndAncestors)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "trunk", now, now, now).
				AddRow("1", "root", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndDescendants)).
		WillReturnError(fmt.Errorf("Task selection failed"))

	taskRepository := NewTask(db)
	err := taskRepository.CreateTaskRelationsBetweenTasks(2, 3)
	if err.Error() != "Task selection failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task selection failed")
	}

	t.Log("Success.")
}

func TestCreateTaskRelationsBetweenTasksFailOnGetAncestorLevel(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID2)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID3)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("3", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("2"))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndAncestors)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "trunk", now, now, now).
				AddRow("1", "root", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndDescendants)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("3", "branch", now, now, now).
				AddRow("4", "leaf", now, now, now))

	// First ancestor loop.
	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnError(fmt.Errorf("level acquisition failed"))

	taskRepository := NewTask(db)
	err := taskRepository.CreateTaskRelationsBetweenTasks(2, 3)
	if err.Error() != "level acquisition failed" {
		t.Fatalf("Got %v\nwant %v", err, "level acquisition failed")
	}

	t.Log("Success.")
}

func TestCreateTaskRelationsBetweenTasksFailOnGetDescendantLevel(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID2)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID3)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("3", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("2"))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndAncestors)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "trunk", now, now, now).
				AddRow("1", "root", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndDescendants)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("3", "branch", now, now, now).
				AddRow("4", "leaf", now, now, now))

	// First ancestor loop.
	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("1"))

	// First descendant loop.
	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnError(fmt.Errorf("level acquisition failed"))

	taskRepository := NewTask(db)
	err := taskRepository.CreateTaskRelationsBetweenTasks(2, 3)
	if err.Error() != "level acquisition failed" {
		t.Fatalf("Got %v\nwant %v", err, "level acquisition failed")
	}

	t.Log("Success.")
}

func TestCreateTaskRelationsBetweenTasksFailOnInsertion(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID2)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectID3)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("3", "テストタスク", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("2"))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndAncestors)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("2", "trunk", now, now, now))

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectSelfAndDescendants)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("3", "branch", now, now, now))

	// First ancestor loop.
	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("1"))

	// First descendant loop.
	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelectPathLength)).
		WillReturnRows(
			sqlmock.NewRows([]string{"path_length"}).
				AddRow("1"))

	// Create section.
	mock.ExpectExec(regexp.QuoteMeta(QueryTaskRelationInsertRegex)).
		WillReturnError(fmt.Errorf("Task relation insertion failed"))

	taskRepository := NewTask(db)
	err := taskRepository.CreateTaskRelationsBetweenTasks(2, 3)
	if err.Error() != "Task relation insertion failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task relation insertion failed")
	}

	t.Log("Success.")
}
