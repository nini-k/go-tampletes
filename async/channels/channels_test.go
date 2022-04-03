package channels

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func wrap(id int) func() error {
	return func() error {
		if id%3 == 0 {
			return fmt.Errorf("some error - id=%d", id)
		}
		return nil
	}
}

func TestChannels(t *testing.T) {
	t.Run("async", func(t *testing.T) {
		var (
			expected = errors.New("some error")
			actual   = <-async[error](func() error {
				return errors.New("some error")
			})
		)

		require.Equal(t, expected, actual)
	})

	t.Run("merge", func(t *testing.T) {
		const N = 6

		functions := make([]func() error, 0, N)
		for i := 1; i <= N; i++ {
			functions = append(functions, wrap(i))
		}

		asyncs := make([]chan error, 0, len(functions))
		for _, f := range functions {
			asyncs = append(asyncs, async[error](f))
		}

		var (
			expected = []error{
				nil,
				nil,
				errors.New("some error - id=3"),
				nil,
				nil,
				errors.New("some error - id=6"),
			}
			actual = make([]error, 0, len(expected))
		)

		for e := range merge[error](asyncs...) {
			actual = append(actual, e)
		}

		require.ElementsMatch(t, expected, actual)
	})
}
