package scim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestError_Error(t *testing.T) {
	e := Error{
		Description: "hello",
		Code:        401,
	}
	require.Equal(t, e.Description, e.Error())
}
