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

func convertIntToAnyTypeSlice(in []int) []any {
	out := make([]any, 0, len(in))
	for _, val := range in {
		out = append(out, val)
	}
	return out
}

func TestLinkedList(t *testing.T) {
	var (
		someList = NewLinkedList()
		values0  = []int{1, 2, 3, 4, 5, 6}
		values1  = reverseSlice(values0)
	)

	t.Run("linked list is empty", func(t *testing.T) {
		var (
			expected = true
			actual   = someList.IsEmpty()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("try remove at head from empty linked list", func(t *testing.T) {
		someList.RemoveAtHead()

		var (
			expected = 0
			actual   = someList.Size()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("add at tail linked list", func(t *testing.T) {
		for _, p := range values0 {
			someList.AddAtTail(p)
		}

		var (
			expected = len(values0)
			actual   = someList.Size()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("convert to slice after adding values at tail", func(t *testing.T) {
		var (
			expected = convertIntToAnyTypeSlice(values0)
			actual   = someList.ConvertToSlice()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("handle for each nodes after adding values at tail", func(t *testing.T) {
		var (
			expected = values0
			actual   = make([]int, 0, len(expected))
		)

		someList.ForEach(func(value any) {
			actual = append(actual, value.(int))
		})

		require.Equal(t, expected, actual)
	})

	t.Run("add at head values into linked list", func(t *testing.T) {
		for _, p := range values0 {
			someList.AddAtHead(p)
		}

		var (
			expected = len(values0) + len(values1)
			actual   = someList.Size()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("convert to slice after adding values at head", func(t *testing.T) {
		var (
			expected = convertIntToAnyTypeSlice(append(values1, values0...))
			actual   = someList.ConvertToSlice()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("handle for each nodes after adding values at head", func(t *testing.T) {
		var (
			expected = append(values1, values0...)
			actual   = make([]int, 0, len(expected))
		)

		someList.ForEach(func(value any) {
			actual = append(actual, value.(int))
		})

		require.Equal(t, expected, actual)
	})

	t.Run("remove at head values from linked list", func(t *testing.T) {
		for range values1 {
			someList.RemoveAtHead()
		}

		var (
			expected = len(values0)
			actual   = someList.Size()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("convert to slice after removing values at head", func(t *testing.T) {
		var (
			expected = convertIntToAnyTypeSlice(values0)
			actual   = someList.ConvertToSlice()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("handle for each nodes after removing values at head", func(t *testing.T) {
		var (
			expected = values0
			actual   = make([]int, 0, len(expected))
		)

		someList.ForEach(func(value any) {
			actual = append(actual, value.(int))
		})

		require.Equal(t, expected, actual)
	})

	t.Run("get head value from linked list", func(t *testing.T) {
		var (
			expected  any = values0[0]
			actual, _     = someList.Head()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("get tail value from linked list", func(t *testing.T) {
		var (
			expected  any = values0[len(values0)-1]
			actual, _     = someList.Tail()
		)

		require.Equal(t, expected, actual)
	})
}
