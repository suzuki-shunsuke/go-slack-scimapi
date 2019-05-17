package scim

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_setPageToQuery(t *testing.T) {
	data := []struct {
		page *Pagination
		exp  url.Values
	}{
		{},
		{
			page: &Pagination{Count: 1, StartIndex: 2},
			exp:  url.Values{"count": []string{"1"}, "startIndex": []string{"2"}},
		},
	}
	for _, d := range data {
		query := url.Values{}
		setPageToQuery(d.page, query)
		if d.exp != nil {
			require.Equal(t, d.exp, query)
		}
	}
}
