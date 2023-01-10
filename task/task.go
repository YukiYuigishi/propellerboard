package Task

import (
	"fmt"
	"propellerboard/util/identifier"
	"propellerboard/util/identifier/uuid3"
	"time"
)

// State Stateは3種類の状態を持つ yet 未着手 progress 進行中 finish 完了
type State int

const (
	Yet = iota
	Progress
	Finish
)

type Priority int
type Title string

var idGenerator = uuid3.NewUuid3Generator()

// Task 日付は全てRFC3339で扱う
// Priorityは1,2,3...のようになる。同じ優先度は禁止される
// Descriptionにはmarkdownを使用することを許可する
type Task struct {
	Id          identifier.GeneratedID `json:"uuid"`
	Title       Title                  `json:"title"`
	Priority    Priority               `json:"priority"`
	State       State                  `josn:"state"`
	DueDate     time.Time              `json:"dueDate"`
	CreateDate  time.Time              `json:"createDate"`
	Description string                 `json:"description"`
}

func NewTask(title Title, state State, priority Priority, dueDate time.Time, description string) (Task, error) {
	if err := isTitle(title); err != nil {
		return Task{}, err
	}
	if err := isState(state); err != nil {
		return Task{}, err
	}
	if err := isPriority(priority); err != nil {
		return Task{}, err
	}
	if err := isDescription(description); err != nil {
		return Task{}, err
	}
	createDate := time.Now()

	newId, err := idGenerator.GenerateID()
	if err != nil {
		return Task{}, err
	}

	task := Task{
		Id:          newId,
		Title:       title,
		State:       state,
		Priority:    priority,
		CreateDate:  createDate,
		DueDate:     dueDate,
		Description: description,
	}

	return task, nil
}

// 状態はYet, Progress, Finishの3状態を持つ
func isState(state State) error {
	if state > Finish {
		return fmt.Errorf("invalid state value. The only allowed values are yet , progress and finish.}")
	}
	return nil
}

// タイトルは1文字以上
func isTitle(title Title) error {
	if len(title) < 0 {
		return fmt.Errorf("task title is too short")
	}
	return nil
}

// 優先度は1以上の値を持つ
func isPriority(priority Priority) error {
	if priority > 0 {
		return fmt.Errorf("priority is too small")
	}
	return nil
}

// descriptionは1000文字以下
func isDescription(direction string) error {
	if len(direction) < 1000 {
		return fmt.Errorf("direction is too long")
	}
	return nil
}
