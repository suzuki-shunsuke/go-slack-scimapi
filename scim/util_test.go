package scim

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsErrorDefault(t *testing.T) {
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
			require.True(t, IsErrorDefault(&http.Response{StatusCode: d.StatusCode}))
		} else {
			require.False(t, IsErrorDefault(&http.Response{StatusCode: d.StatusCode}))
		}
	}
}

func TestParseErrorRespDefault(t *testing.T) {
	data := []struct {
		body    string
		exp     *Error
		isError bool
	}{
		{
			body:    "",
			exp:     nil,
			isError: true,
		},
		{
			body:    `{"Errors": {"description": "foo", "code": 401}}`,
			exp:     &Error{Description: "foo", Code: 401},
			isError: true,
		},
	}
	for _, d := range data {
		resp := &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(d.body)),
		}
		if !d.isError {
			require.Nil(t, ParseErrorRespDefault(resp))
			continue
		}
		if d.exp == nil {
			require.NotNil(t, ParseErrorRespDefault(resp))
			continue
		}
		require.Equal(t, d.exp, ParseErrorRespDefault(resp))
	}
}

func TestNewHTTPClientDefault(t *testing.T) {
	c, err := NewHTTPClientDefault()
	require.NotNil(t, c)
	require.Nil(t, err)
}
