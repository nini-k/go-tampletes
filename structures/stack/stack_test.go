package stack

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

func convertIntToAnyTypeSlice(in []int) []interface{} {
	out := make([]interface{}, 0, len(in))
	for _, val := range in {
		out = append(out, val)
	}
	return out
}

func TestStack(t *testing.T) {
	var (
		stack          = New()
		values         = []int{1, 2, 3, 4, 5, 6}
		reversedValues = reverseSlice(values)
	)

	t.Run("stack is empty", func(t *testing.T) {
		var (
			expected = true
			actual   = stack.IsEmpty()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("try pop from empty stack", func(t *testing.T) {
		stack.Pop()

		var (
			expected = 0
			actual   = stack.Size()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("push values", func(t *testing.T) {
		for _, p := range values {
			stack.Push(p)
		}

		var (
			expected = len(values)
			actual   = stack.Size()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("convert to slice after adding values", func(t *testing.T) {
		var (
			expected = convertIntToAnyTypeSlice(reversedValues)
			actual   = stack.ConvertToSlice()
		)

		require.Equal(t, expected, actual)
	})

	t.Run("pop values", func(t *testing.T) {
		var (
			expected = convertIntToAnyTypeSlice(reversedValues)
			actual   = make([]interface{}, 0, len(expected))
		)

		for !stack.IsEmpty() {
			actual = append(actual, stack.Pop())
		}

		require.Equal(t, expected, actual)
	})
}
