package repository

import (
	"fmt"
	"golang.org/x/xerrors"
	"strings"
	"time"
)

// This file is created to focus on closure table methods of TaskRepository.
// Base methods are placed in `app/domain/repository/task.go`.

type TreeableTask struct {
	Task
	Depth    uint
	Children []TreeableTask
}

type TaskWithDepth struct {
	Task
	Depth uint
}

func (t *TaskRepository) ListTree() (*[]TreeableTask, error) {
	var taskWithDepths []TaskWithDepth

	// The result of this query is ordered by materialized path like `3,5,6,12`
	// WANTFIX:
	// id sort works like below,
	// ```
	// ID: 1, 10, 100, 2, 20, 200, ...
	// ```
	// because id is handled as string in this rule.
	// root level sorting should be done with id as integer.
	query := `
		SELECT
			tasks.id,
			tasks.title,
			tasks.type,
			tasks.completed_at,
			tasks.created_at,
			tasks.updated_at,
			max(descendant_relations.path_length) AS depth
		FROM
			tasks
		LEFT JOIN task_relations AS descendant_relations ON tasks.id = descendant_relations.descendant_id
		GROUP BY
			tasks.id,
			tasks.title,
			tasks.type,
			tasks.completed_at,
			tasks.created_at,
			tasks.updated_at,
			descendant_relations.descendant_id
		ORDER BY
			group_concat(descendant_relations.ancestor_id ORDER BY descendant_relations.path_length DESC), tasks.id
	`

	if err := t.DB.Raw(query).Scan(&taskWithDepths).Error; err != nil {
		return nil, err
	}

	var treeableTask TreeableTask
	treeableTask, _ = t.appendChildren(treeableTask, taskWithDepths, 0)

	return &treeableTask.Children, nil
}

func (t *TaskRepository) FindTreeByID(id uint) (*TreeableTask, error) {
	var taskWithDepths []TaskWithDepth

	// The result of this query is ordered by materialized path like `3,5,6,12`
	// WANTFIX:
	// id sort works like below,
	// ```
	// ID: 1, 10, 100, 2, 20, 200, ...
	// ```
	// because id is handled as string in this rule.
	// root level sorting should be done with id as integer.
	// TODO: do not use IN query. This may cause slow query after db record increased.
	query := `
		SELECT
			tasks.id,
			tasks.title,
			tasks.type,
			tasks.completed_at,
			tasks.created_at,
			tasks.updated_at,
			max(descendant_relations.path_length) AS depth
		FROM
			tasks
		LEFT JOIN task_relations AS descendant_relations ON tasks.id = descendant_relations.descendant_id
		WHERE (
			tasks.id IN (
				SELECT
					tasks.id
				FROM
					tasks
				LEFT JOIN task_relations ON tasks.id = task_relations.descendant_id
				WHERE (
					task_relations.ancestor_id = ?
				)
			)
		)
		GROUP BY
			tasks.id,
			tasks.title,
			tasks.type,
			tasks.completed_at,
			tasks.created_at,
			tasks.updated_at
		ORDER BY
			group_concat(descendant_relations.ancestor_id ORDER BY descendant_relations.path_length DESC), tasks.id
	`

	if err := t.DB.Raw(query, id).Scan(&taskWithDepths).Error; err != nil {
		return nil, err
	}

	var treeableTask TreeableTask
	treeableTask, _ = t.appendChildren(treeableTask, taskWithDepths, 0)
	treeableTask = treeableTask.Children[0]

	return &treeableTask, nil
}

func (t *TaskRepository) appendChildren(tree TreeableTask, tasks []TaskWithDepth, idx uint) (TreeableTask, uint) {
	if uint(len(tasks)) <= idx {
		return tree, idx
	}

	task := tasks[idx]
	if tree.Depth < task.Depth {
		child := TreeableTask{
			Task: Task{
				ID:          task.ID,
				Title:       task.Title,
				Type:        task.Type,
				CompletedAt: task.CompletedAt,
				CreatedAt:   task.CreatedAt,
				UpdatedAt:   task.UpdatedAt,
			},
			Depth: task.Depth,
		}

		child, idx = t.appendChildren(child, tasks, idx+1)
		tree.Children = append(tree.Children, child)
		return t.appendChildren(tree, tasks, idx)
	}

	return tree, idx
}

func (t *TaskRepository) ListSelfAndDescendants(taskID uint) (*[]Task, error) {
	var tasks []Task

	query := `
		SELECT
			tasks.id,
			tasks.title,
			tasks.type,
			tasks.completed_at,
			tasks.created_at,
			tasks.updated_at
		FROM tasks
		LEFT JOIN task_relations AS descendant_relations ON tasks.id = descendant_relations.descendant_id
		WHERE
			descendant_relations.ancestor_id = ?
		GROUP BY
			tasks.id,
			tasks.title,
			tasks.type,
			tasks.completed_at,
			tasks.created_at,
			tasks.updated_at,
			descendant_relations.path_length
		ORDER BY descendant_relations.path_length asc
	`

	if err := t.DB.Raw(query, taskID).Scan(&tasks).Error; err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		// Imitating gorm `record not found` error object.
		// This should be replaced with RecordNotFound error object.
		return nil, xerrors.New("record not found")
	}

	return &tasks, nil
}

func (t *TaskRepository) ListSelfAndAncestors(taskID uint) (*[]Task, error) {
	var tasks []Task

	query := `
		SELECT
			tasks.id,
			tasks.title,
			tasks.type,
			tasks.completed_at,
			tasks.created_at,
			tasks.updated_at
		FROM tasks
		LEFT JOIN task_relations AS ancestor_relations ON tasks.id = ancestor_relations.ancestor_id
		WHERE
			ancestor_relations.descendant_id = ?
		GROUP BY
			tasks.id,
			tasks.title,
			tasks.type,
			tasks.completed_at,
			tasks.created_at,
			tasks.updated_at,
			ancestor_relations.path_length
		ORDER BY ancestor_relations.path_length asc
	`

	if err := t.DB.Raw(query, taskID).Scan(&tasks).Error; err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		// Imitating gorm `record not found` error object.
		// This should be replaced with RecordNotFound error object.
		return nil, xerrors.New("record not found")
	}

	return &tasks, nil
}

func (t *TaskRepository) GetLevel(taskID uint) (uint, error) {
	level := struct{ PathLength uint }{}

	query := `
		SELECT max(path_length) AS path_length
		FROM task_relations
		WHERE descendant_id = ?
	`
	if err := t.DB.Raw(query, taskID).Scan(&level).Error; err != nil {
		return 0, err
	}

	return level.PathLength, nil
}

func (t *TaskRepository) DeleteAncestorTaskRelations(taskID uint) error {
	descendantTasks, err := t.ListSelfAndDescendants(taskID)
	if err != nil {
		return err
	}

	var descendantTaskIDs []uint
	for _, descendantTask := range *descendantTasks {
		descendantTaskIDs = append(descendantTaskIDs, descendantTask.ID)
	}

	query := `
		DELETE FROM task_relations
		WHERE descendant_id IN (?)
		AND ancestor_id NOT IN (?)
	`

	if err := t.DB.Exec(query, descendantTaskIDs, descendantTaskIDs).Error; err != nil {
		return err
	}

	return nil
}

func (t *TaskRepository) CreateTaskRelationsBetweenTasks(parentTaskID uint, childTaskID uint) error {
	var taskRelations []TaskRelation

	var parent Task
	if err := t.DB.Find(&parent, parentTaskID).Error; err != nil {
		return err
	}

	var child Task
	if err := t.DB.Find(&child, childTaskID).Error; err != nil {
		return err
	}

	parentLevel, err := t.GetLevel(parent.ID)
	if err != nil {
		return err
	}

	ancestors, err := t.ListSelfAndAncestors(parent.ID)
	if err != nil {
		return err
	}

	descendants, err := t.ListSelfAndDescendants(child.ID)
	if err != nil {
		return err
	}

	for _, ancestor := range *ancestors {
		ancestorLevel, err := t.GetLevel(ancestor.ID)
		if err != nil {
			return err
		}
		pathLength := parentLevel - ancestorLevel + 1
		for _, descendant := range *descendants {
			descendantLevel, err := t.GetLevel(descendant.ID)
			if err != nil {
				return err
			}
			descendantPathLength := pathLength + descendantLevel
			taskRelation := TaskRelation{
				AncestorID:   ancestor.ID,
				DescendantID: descendant.ID,
				PathLength:   descendantPathLength,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}
			taskRelations = append(taskRelations, taskRelation)
		}
	}

	fmt.Printf("%v", taskRelations)

	// Note that after `VALUES` there is one space.
	query := `
		INSERT INTO task_relations (
			ancestor_id,
			descendant_id,
			path_length,
			created_at,
			updated_at
		) VALUES 
	`
	var values []string
	for _, taskRelation := range taskRelations {
		value := fmt.Sprintf(
			"(%d, %d, %d, '%s', '%s')",
			taskRelation.AncestorID,
			taskRelation.DescendantID,
			taskRelation.PathLength,
			taskRelation.CreatedAt.Format("2006-01-02 15:04:05"),
			taskRelation.UpdatedAt.Format("2006-01-02 15:04:05"),
		)
		values = append(values, value)
	}
	query += strings.Join(values, ", ")
	if err := t.DB.Exec(query).Error; err != nil {
		return err
	}

	return nil
}
