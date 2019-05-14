package scim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	require.NotNil(t, NewClient("XXX"))
}
