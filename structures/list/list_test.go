package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func reverseSlice(in []int) []int {
	out := make([]int, len(in))
	r, l := 0, len(in)-1
	for r <= l {
		out[l], out[r] = in[r], in[l]
		l--
		r++
	}
	return out
}

func TestLinkedList(t *testing.T) {
	var (
		list           = New[int]()
		values         = []int{1, 2, 3, 4, 5, 6}
		reversedValues = reverseSlice(values)
	)

	t.Run("linked list is empty", func(t *testing.T) {
		var (
			expected = true
			actual   = list.IsEmpty()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("try remove at head from empty linked list", func(t *testing.T) {
		list.RemoveAtHead()

		var (
			expected = 0
			actual   = list.Size()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("add at tail linked list", func(t *testing.T) {
		for _, p := range values {
			list.AddAtTail(p)
		}

		var (
			expected = len(values)
			actual   = list.Size()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("convert to slice after adding values at tail", func(t *testing.T) {
		var (
			expected = values
			actual   = list.ConvertToSlice()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("handle for each nodes after adding values at tail", func(t *testing.T) {
		var (
			expected = values
			actual   = make([]int, 0, len(expected))
		)

		list.ForEach(func(value int) {
			actual = append(actual, value)
		})

		require.Equal(t, expected, actual)
	})

	t.Run("add at head values into linked list", func(t *testing.T) {
		for _, p := range values {
			list.AddAtHead(p)
		}

		var (
			expected = len(values) + len(reversedValues)
			actual   = list.Size()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("convert to slice after adding values at head", func(t *testing.T) {
		var (
			expected = append(reversedValues, values...)
			actual   = list.ConvertToSlice()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("handle for each nodes after adding values at head", func(t *testing.T) {
		var (
			expected = append(reversedValues, values...)
			actual   = make([]int, 0, len(expected))
		)

		list.ForEach(func(value int) {
			actual = append(actual, value)
		})

		require.Equal(t, expected, actual)
	})

	t.Run("remove at head values from linked list", func(t *testing.T) {
		for range reversedValues {
			list.RemoveAtHead()
		}

		var (
			expected = len(values)
			actual   = list.Size()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("convert to slice after removing values at head", func(t *testing.T) {
		var (
			expected = values
			actual   = list.ConvertToSlice()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("handle for each nodes after removing values at head", func(t *testing.T) {
		var (
			expected = values
			actual   = make([]int, 0, len(expected))
		)

		list.ForEach(func(value int) {
			actual = append(actual, value)
		})

		require.Equal(t, expected, actual)
	})

	t.Run("get head value from linked list", func(t *testing.T) {
		var (
			expected  interface{} = values[0]
			actual, _             = list.Head()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("get tail value from linked list", func(t *testing.T) {
		var (
			expected  interface{} = values[len(values)-1]
			actual, _             = list.Tail()
		)

		require.Equal(t, expected, actual)
	})
}
