package scim

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isError(t *testing.T) {
	data := []struct {
		StatusCode int
		isError    bool
	}{
		{
			StatusCode: 400,
			isError:    true,
		},
		{
			StatusCode: 500,
			isError:    true,
		},
		{
			StatusCode: 200,
			isError:    false,
		},
		{
			StatusCode: 399,
			isError:    false,
		},
	}
	for _, d := range data {
		if d.isError {
			require.True(t, isError(&http.Response{StatusCode: d.StatusCode}))
		} else {
			require.False(t, isError(&http.Response{StatusCode: d.StatusCode}))
		}
	}
}

func Test_clientFn(t *testing.T) {
	c, err := clientFn()
	require.NotNil(t, c)
	require.Nil(t, err)
}
