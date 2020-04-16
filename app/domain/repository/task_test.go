package repository

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	QueryTaskInsert            = `INSERT INTO "tasks" ("title","type","completed_at","starts_at","expires_at","created_at","updated_at") VALUES (?,?,?,?,?,?,?)`
	QueryTaskRelationInsert    = `INSERT INTO "task_relations" ("ancestor_id","descendant_id","path_length","created_at","updated_at") VALUES (?,?,?,?,?)`
	QueryTaskSelect            = `SELECT * FROM "tasks" WHERE ("tasks"."id" = 1)`
	QueryTaskUpdateCompletedAt = `UPDATE "tasks" SET "completed_at" = ?`
	QueryTaskUpdateType        = `UPDATE "tasks" SET "type" = ?`
)

func TestCreate(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(QueryTaskInsert)).
		WillReturnResult(sqlmock.NewResult(2, 1))

	mock.ExpectExec(regexp.QuoteMeta(QueryTaskRelationInsert)).
		WillReturnResult(sqlmock.NewResult(2, 1))

	mock.ExpectCommit()

	taskRepository := NewTask(db)
	title := "新しいテストタスク"
	_, err := taskRepository.Create(title)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestCreateWithoutTitle(t *testing.T) {
	db, _, _ := getDBMock()
	defer db.Close()

	taskRepository := NewTask(db)

	// Set blank title.
	title := ""
	_, err := taskRepository.Create(title)
	if err.Error() != "title can't have blank value" {
		t.Fatalf("Got %v\nwant %v", err, "title can't have blank value")
	}

	t.Log("Success.")
}

func TestCreateRollbackByTaskInsertion(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(QueryTaskInsert)).
		WillReturnError(fmt.Errorf("Task insertion failed"))

	mock.ExpectRollback()

	taskRepository := NewTask(db)
	title := "新しいテストタスク"
	_, err := taskRepository.Create(title)
	if err.Error() != "Task insertion failed" {
		t.Fatalf("Got %v\nwant %v", err, "Task insertion failed")
	}

	t.Log("Success.")
}

func TestCreateRollbackByTaskRelationInsertion(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(QueryTaskInsert)).
		WillReturnResult(sqlmock.NewResult(2, 1))

	mock.ExpectExec(regexp.QuoteMeta(QueryTaskRelationInsert)).
		WillReturnError(fmt.Errorf("TaskRelation insertion failed"))

	mock.ExpectRollback()

	taskRepository := NewTask(db)
	title := "新しいテストタスク"
	_, err := taskRepository.Create(title)
	if err.Error() != "TaskRelation insertion failed" {
		t.Fatalf("Got %v\nwant %v", err, "TaskRelation insertion failed")
	}

	t.Log("Success.")
}

func TestUpdateCompletedAt(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelect)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}).
				AddRow("1", "テストタスク", nil, now, now))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(QueryTaskUpdateCompletedAt)).
		WithArgs(now, AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	taskRepository := NewTask(db)

	err := taskRepository.UpdateCompletedAt(1, now)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestUpdateCompletedAtWithNotFoundID(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelect)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "completed_at", "created_at", "updated_at"}))

	taskRepository := NewTask(db)

	err := taskRepository.UpdateCompletedAt(1, now)

	if err.Error() != "record not found" {
		t.Fatalf("Got %v\nwant %v", err, "record not found")
	}

	t.Log("Success.")
}

func TestUpdateType(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelect)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "type", "completed_at", "created_at", "updated_at"}).
				AddRow("1", "テストタスク", 0, nil, now, now))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(QueryTaskUpdateType)).
		WithArgs(10, AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	taskRepository := NewTask(db)

	err := taskRepository.UpdateType(1, 10)

	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestUpdateTypeSelectionCallsError(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelect)).
		WillReturnError(fmt.Errorf("task selection failed"))

	taskRepository := NewTask(db)

	err := taskRepository.UpdateType(1, 10)

	if err.Error() != "task selection failed" {
		t.Fatalf("Got %v\nwant %v", err, "record not found")
	}

	t.Log("Success.")
}

func TestUpdateTypeUpdateCallsError(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(QueryTaskSelect)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "type", "completed_at", "created_at", "updated_at"}).
				AddRow("1", "テストタスク", 0, nil, now, now))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(QueryTaskUpdateType)).
		WillReturnError(fmt.Errorf("updating task failed"))
	mock.ExpectCommit()

	taskRepository := NewTask(db)

	err := taskRepository.UpdateType(1, 10)

	if err.Error() != "updating task failed" {
		t.Fatalf("Got %v\nwant %v", err, "record not found")
	}

	t.Log("Success.")
}
