package Task

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import "time"

func TestCreateTask(t *testing.T) {

	// 正常系
	t.Run("Check normal value", func(t *testing.T) {
		title := Title("test")
		state := State(Yet)
		priority := Priority(1)
		dueTime := time.Date(2022, 12, 18, 0, 0, 0, 0, time.UTC)
		description := ""

		newTask, err := NewTask(title, state, priority, dueTime, description)
		assert.NoError(t, err)

		// 値が正常か確認
		assert.Equal(t, title, newTask.Title)
		assert.Equal(t, state, newTask.State)
		assert.Equal(t, priority, newTask.Priority)
		assert.Equal(t, dueTime, newTask.DueDate)
		assert.Equal(t, description, newTask.Description)
	})
	// 異常系
	t.Run("Check Title error", func(t *testing.T) {
		title := Title("")
		state := State(Yet)
		priority := Priority(1)
		dueTime := time.Date(2022, 12, 18, 0, 0, 0, 0, time.UTC)
		description := ""

		_, err := NewTask(title, state, priority, dueTime, description)

		assert.Error(t, err)
	})
	// 異常系
	t.Run("Check state error", func(t *testing.T) {
		title := Title("test")
		state := State(3)
		priority := Priority(1)
		dueTime := time.Date(2022, 12, 18, 0, 0, 0, 0, time.UTC)
		description := ""

		_, err := NewTask(title, state, priority, dueTime, description)

		assert.Error(t, err)
	})
	// 異常系
	t.Run("Check priority error", func(t *testing.T) {
		title := Title("test")
		state := State(Yet)
		priority := Priority(0)
		dueTime := time.Date(2022, 12, 18, 0, 0, 0, 0, time.UTC)
		description := ""

		_, err := NewTask(title, state, priority, dueTime, description)

		assert.Error(t, err)
	})
	// 異常系
	t.Run("Check description error", func(t *testing.T) {
		title := Title("test")
		state := State(Yet)
		priority := Priority(1)
		dueTime := time.Date(2022, 12, 18, 0, 0, 0, 0, time.UTC)
		description := ""
		for i := 0; i < 101; i++ {
			description += "0123456789"
		}

		_, err := NewTask(title, state, priority, dueTime, description)

		assert.Error(t, err)
	})

}
