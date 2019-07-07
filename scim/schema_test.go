package scim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchema_UnmarshalJSON(t *testing.T) {
	data := []struct {
		body    string
		isError bool
		exp     *Schema
	}{
		{
			body:    "",
			isError: true,
		},
		{
			body:    "{}",
			isError: false,
			exp:     &Schema{},
		},
		{
			body:    `{"schema": 0}`,
			isError: true,
		},
	}

	for _, d := range data {
		schema := &Schema{}
		err := schema.UnmarshalJSON([]byte(d.body))
		if d.isError {
			require.NotNil(t, err)
			continue
		}
		require.Nil(t, err)
		if d.exp != nil {
			require.Equal(t, d.exp, schema)
		}
	}
}

func TestAttribute_UnmarshalJSON(t *testing.T) {
	data := []struct {
		body    string
		isError bool
		exp     *Schema
	}{
		{
			body:    "",
			isError: true,
		},
		{
			body:    `{"subAttributes": 0}`,
			isError: true,
		},
	}

	for _, d := range data {
		attr := &Attribute{}
		err := attr.UnmarshalJSON([]byte(d.body))
		if d.isError {
			require.NotNil(t, err)
			continue
		}
		require.Nil(t, err)
		if d.exp != nil {
			require.Equal(t, d.exp, attr)
		}
	}
}
