package repository

import (
	"time"

	"golang.org/x/xerrors"

	"github.com/jinzhu/gorm"
)

// TaskRepository is struct with DB connection.
// See also `app/domain/repository/task_closure_table.go`
type TaskRepository struct {
	DB *gorm.DB
}

// Task is struct for gorm mapping.
type Task struct {
	ID          uint       `gorm:"not null;index"`
	Title       string     `gorm:"not null;index"`
	Type        uint       `gorm:"not null"`
	CompletedAt *time.Time `sql:"type:datetime"`
	ExpiresAt   *time.Time `sql:"type:datetime"`
	CreatedAt   time.Time  `gorm:"not null" sql:"type:datetime"`
	UpdatedAt   time.Time  `gorm:"not null" sql:"type:datetime"`
}

// TaskRelation is struct for gorm mapping.
type TaskRelation struct {
	ID           uint      `gorm:"not null;index"`
	AncestorID   uint      `gorm:"not null;index"`
	DescendantID uint      `gorm:"not null;index"`
	PathLength   uint      `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null" sql:"type:datetime"`
	UpdatedAt    time.Time `gorm:"not null" sql:"type:datetime"`
}

// NewTask returns TaskRepository with DB connection.
func NewTask(db *gorm.DB) *TaskRepository {
	var t TaskRepository
	t.DB = db
	return &t
}

// Create saves a task record through persistence layer.
func (t *TaskRepository) Create(title string) (*Task, error) {
	if title == "" {
		return nil, xerrors.New("title can't have blank value")
	}

	tx := t.DB.Begin()

	task := Task{Title: title, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := tx.Create(&task).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// This process is written to make closure table work.
	taskRelation := TaskRelation{AncestorID: task.ID, DescendantID: task.ID, PathLength: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := tx.Create(&taskRelation).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &task, nil
}

// UpdateCompletedAt updates completed_at in a task record.
func (t *TaskRepository) UpdateCompletedAt(id uint, completedAt time.Time) error {
	var task Task
	if err := t.DB.Find(&task, id).Error; err != nil {
		return err
	}

	// This won't happen expected error.
	t.DB.Model(&task).Update("CompletedAt", completedAt)
	return nil
}

// UpdateTitle updates title in a task record.
func (t *TaskRepository) UpdateTitle(id uint, title string) error {
	var task Task
	if err := t.DB.Find(&task, id).Error; err != nil {
		return err
	}

	if err := t.DB.Model(&task).Update("title", title).Error; err != nil {
		return err
	}
	return nil
}

// UpdateType updates type in a task record.
func (t *TaskRepository) UpdateType(id uint, taskType uint) error {
	var task Task
	if err := t.DB.Find(&task, id).Error; err != nil {
		return err
	}

	if err := t.DB.Model(&task).Update("Type", taskType).Error; err != nil {
		return err
	}

	return nil
}
