package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/whalepod/milelane/app/infrastructure"
)

const (
	QueryTaskTreeSelect = `SELECT tasks.id, tasks.title, tasks.type, tasks.completed_at, tasks.created_at, tasks.updated_at, max(descendant_relations.path_length) AS depth FROM tasks LEFT JOIN task_relations AS descendant_relations ON tasks.id = descendant_relations.descendant_id GROUP BY tasks.id, tasks.title, tasks.type, tasks.completed_at, tasks.created_at, tasks.updated_at, descendant_relations.descendant_id ORDER BY group_concat(descendant_relations.ancestor_id ORDER BY descendant_relations.path_length DESC), tasks.id`
)

func TestTaskIndex(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.GET("/tasks", func(c *gin.Context) {
		TaskIndex(c)
	})

	req, _ := http.NewRequest("GET", "/tasks", nil)
	r.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Fatal("Returned wrong http status.")
	}

	t.Log("Success.")
}

func TestTaskIndexFailByInfrastructure(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.GET("/tasks", func(c *gin.Context) {
		TaskIndex(c)
	})

	// In case infrastructure.DB broken, it returns StatusInternalServerError.
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(QueryTaskTreeSelect)).
		WillReturnError(fmt.Errorf("Task insertion failed"))

	// Mock infrastructure.DB to test irregular error.
	originalDB := infrastructure.DB
	infrastructure.DB = db

	req, _ := http.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusInternalServerError != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	infrastructure.DB = originalDB
	t.Log("Success.")
}

func TestTaskIndexWithDeviceUUID(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.GET("/tasks", func(c *gin.Context) {
		TaskIndex(c)
	})

	req, _ := http.NewRequest("GET", "/tasks", nil)
	req.Header.Set("X-Device-UUID", "60982a48-9328-441f-805b-d3ab0cad9e1f")
	r.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskCreate(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks", func(c *gin.Context) {
		TaskCreate(c)
	})

	// With valid title, it returns StatusOK.
	jsonStr := `{"title":"テストタイトル"}`
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	expectedBodyPart := "\"title\":\"テストタイトル\""
	if !strings.Contains(res.Body.String(), expectedBodyPart) {
		t.Fatalf("Returned wrong http body. Actual body: %v, Expected to have %v", res.Body.String(), expectedBodyPart)
	}

	t.Log("Success.")
}

func TestTaskCreateWithDeviceUUID(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks", func(c *gin.Context) {
		TaskCreate(c)
	})

	// With valid title, it returns StatusOK.
	jsonStr := `{"title":"テストタイトル"}`
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device-UUID", "60982a48-9328-441f-805b-d3ab0cad9e1f")
	r.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	expectedBodyPart := "\"title\":\"テストタイトル\""
	if !strings.Contains(res.Body.String(), expectedBodyPart) {
		t.Fatalf("Returned wrong http body. Actual body: %v, Expected to have %v", res.Body.String(), expectedBodyPart)
	}

	t.Log("Success.")
}

func TestTaskCreateWithVacantTitle(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks", func(c *gin.Context) {
		TaskCreate(c)
	})

	// With wrong title, it returns StatusUnprocessableEntity.
	jsonStr := `{"title":""}`
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusUnprocessableEntity != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskCreateWithoutTitle(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks", func(c *gin.Context) {
		TaskCreate(c)
	})

	// Without title key, it returns StatusUnprocessableEntity.
	req, _ := http.NewRequest("POST", "/tasks", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusUnprocessableEntity != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskCreateFailByInfrastructure(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks", func(c *gin.Context) {
		TaskCreate(c)
	})

	// In case infrastructure.DB broken, it returns StatusInternalServerError.
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "tasks" ("title","completed_at","expires_at","created_at","updated_at") VALUES (?,?,?,?,?)`)).
		WillReturnError(fmt.Errorf("Task insertion failed"))

	// Mock infrastructure.DB to test irregular error.
	originalDB := infrastructure.DB
	infrastructure.DB = db

	jsonStr := `{"title":"テストタイトル"}`
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusInternalServerError != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	infrastructure.DB = originalDB
	t.Log("Success.")
}

func TestTaskShow(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.GET("/tasks/:taskID", func(c *gin.Context) {
		TaskShow(c)
	})

	// With valid taskID, it returns StatusOK.
	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	r.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskShowWithNotFoundID(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.GET("/tasks/:taskID", func(c *gin.Context) {
		TaskShow(c)
	})

	// With wrong taskID, it returns StatusNotFound.
	req, _ := http.NewRequest("GET", "/tasks/9999", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusNotFound != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskShowWithInvalidPath(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.GET("/tasks/:taskID", func(c *gin.Context) {
		TaskShow(c)
	})

	// With wrong taskID, it returns StatusNotFound.
	req, _ := http.NewRequest("GET", "/tasks/wrong_path", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusBadRequest != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskComplete(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/complete", func(c *gin.Context) {
		TaskComplete(c)
	})

	// With valid taskID, it returns StatusOK.
	req, _ := http.NewRequest("POST", "/tasks/1/complete", nil)
	r.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskCompleteWithNotFoundID(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/complete", func(c *gin.Context) {
		TaskComplete(c)
	})

	// With wrong taskID, it returns StatusNotFound.
	req, _ := http.NewRequest("POST", "/tasks/9999/complete", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusNotFound != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskCompleteWithInvalidPath(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/complete", func(c *gin.Context) {
		TaskComplete(c)
	})

	// With wrong taskID, it returns StatusNotFound.
	req, _ := http.NewRequest("POST", "/tasks/wrong_path/complete", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusBadRequest != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskLanize(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/lanize", func(c *gin.Context) {
		TaskLanize(c)
	})

	// With valid taskID, it returns StatusOK.
	req, _ := http.NewRequest("POST", "/tasks/1/lanize", nil)
	r.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskLanizeWithNotFoundID(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/lanize", func(c *gin.Context) {
		TaskLanize(c)
	})

	// With wrong taskID, it returns StatusNotFound.
	req, _ := http.NewRequest("POST", "/tasks/9999/lanize", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusNotFound != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskLanizeWithInvalidPath(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/lanize", func(c *gin.Context) {
		TaskLanize(c)
	})

	// With wrong taskID, it returns StatusNotFound.
	req, _ := http.NewRequest("POST", "/tasks/wrong_path/lanize", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusBadRequest != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskMoveToRoot(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/move-to-root", func(c *gin.Context) {
		TaskMoveToRoot(c)
	})

	// With valid taskID, it returns StatusOK.
	req, _ := http.NewRequest("POST", "/tasks/1/move-to-root", nil)
	r.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskMoveToRootWithNotFoundID(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/move-to-root", func(c *gin.Context) {
		TaskMoveToRoot(c)
	})

	// With wrong taskID, it returns StatusNotFound.
	req, _ := http.NewRequest("POST", "/tasks/9999/move-to-root", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusNotFound != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskMoveToRootWithInvalidPath(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/move-to-root", func(c *gin.Context) {
		TaskMoveToRoot(c)
	})

	// With wrong taskID, it returns StatusNotFound.
	req, _ := http.NewRequest("POST", "/tasks/wrong_path/move-to-root", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusBadRequest != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskMoveToChild(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/move-to-child/:parentTaskID", func(c *gin.Context) {
		TaskMoveToChild(c)
	})

	// With valid taskID, it returns StatusOK.
	req, _ := http.NewRequest("POST", "/tasks/1/move-to-child/2", nil)
	r.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskMoveToChildWithNotFoundTaskID(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/move-to-child/:parentTaskID", func(c *gin.Context) {
		TaskMoveToChild(c)
	})

	req, _ := http.NewRequest("POST", "/tasks/9999/move-to-child/2", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusNotFound != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskMoveToChildWithInvalidTaskPath(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/move-to-child/:parentTaskID", func(c *gin.Context) {
		TaskMoveToChild(c)
	})

	req, _ := http.NewRequest("POST", "/tasks/wrong_path/move-to-child/2", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusBadRequest != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskMoveToChildWithNotFoundParentTaskID(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/move-to-child/:parentTaskID", func(c *gin.Context) {
		TaskMoveToChild(c)
	})

	req, _ := http.NewRequest("POST", "/tasks/1/move-to-child/9999", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusNotFound != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestTaskMoveToChildWithInvalidParentTaskPath(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/tasks/:taskID/move-to-child/:parentTaskID", func(c *gin.Context) {
		TaskMoveToChild(c)
	})

	req, _ := http.NewRequest("POST", "/tasks/1/move-to-child/wrong_path", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusBadRequest != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}
