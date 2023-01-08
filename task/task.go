package Task

import (
	"fmt"
	"time"
   "propellerboard/util/identifier"
   "propellerboard/util/identifier/uuid3"
)

//{
//    Task
//}

// Stateは3種類の状態を持つ
// yet 未着手
// progress 進行中
// finish 完了
type State int
idGenerator := 
const (
	Yet = iota
	Progress
	Finish
)

// 日付は全てRFC3339で扱う
// Priorityは1,2,3...のようになる。同じ優先度は禁止される
// Descriptionにはmarkdownを使用することを許可する
type Task struct {
	Uuid        string    `json:"uuid"`
	Title       string    `json:"title"`
	Priority    int       `json:"priority"`
	State       State     `josn:"state"`
	DueDate     time.Time `json:"dueDate"`
	CreateDate  time.Time `json:"createDate"`
	Description string    `json:"description"`
}

func NewTask(title string, state State, dueDate time.Time, createDate time.Time, description string) (Task, error) {
	if err := isState(state); err != nil {
		return Task{}, err
	}
	task := Task{
   }

	return task, nil
}

func isState(state State) error {
	if state > Finish {
		return fmt.Errorf("Invalid state value. The only allowed values are yet , progress and finish.}")
	}

	return nil
}
