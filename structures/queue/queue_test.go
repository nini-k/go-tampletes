package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	var (
		queue  = New[int]()
		values = []int{1, 2, 3, 4, 5, 6}
	)

	t.Run("queue is empty", func(t *testing.T) {
		var (
			expected = true
			actual   = queue.IsEmpty()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("try pop from empty queue", func(t *testing.T) {
		queue.Pop()

		var (
			expected = 0
			actual   = queue.Size()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("push values", func(t *testing.T) {
		for _, p := range values {
			queue.Push(p)
		}

		var (
			expected = len(values)
			actual   = queue.Size()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("convert to slice after adding values", func(t *testing.T) {
		var (
			expected = values
			actual   = queue.ConvertToSlice()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("pop values", func(t *testing.T) {
		var (
			expected = values
			actual   = make([]int, 0, len(expected))
		)

		for !queue.IsEmpty() {
			actual = append(actual, queue.Pop())
		}

		require.Equal(t, expected, actual)
	})
}
