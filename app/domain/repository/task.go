package repository

import (
	"golang.org/x/xerrors"
	"time"

	"github.com/jinzhu/gorm"
)

// See also `app/domain/repository/task_closure_table.go`
type TaskRepository struct {
	DB *gorm.DB
}

// Struct for gorm mapping.
type Task struct {
	ID          uint       `gorm:"not null;index"`
	Title       string     `gorm:"not null;index"`
	Type        uint       `gorm:"not null"`
	CompletedAt *time.Time `sql:"type:datetime"`
	CreatedAt   time.Time  `gorm:"not null" sql:"type:datetime"`
	UpdatedAt   time.Time  `gorm:"not null" sql:"type:datetime"`
}

// Struct for gorm mapping.
type TaskRelation struct {
	ID           uint      `gorm:"not null;index"`
	AncestorID   uint      `gorm:"not null;index"`
	DescendantID uint      `gorm:"not null;index"`
	PathLength   uint      `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null" sql:"type:datetime"`
	UpdatedAt    time.Time `gorm:"not null" sql:"type:datetime"`
}

func NewTask(db *gorm.DB) *TaskRepository {
	var t TaskRepository
	t.DB = db
	return &t
}

func (t *TaskRepository) Create(title string) (*Task, error) {
	if title == "" {
		return nil, xerrors.New("Title can't have blank value.")
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

func (t *TaskRepository) UpdateCompletedAt(id uint, completedAt time.Time) error {
	var task Task
	if err := t.DB.Find(&task, id).Error; err != nil {
		return err
	}

	// This won't happen expected error.
	t.DB.Model(&task).Update("CompletedAt", completedAt)
	return nil
}

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
