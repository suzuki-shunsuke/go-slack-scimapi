package scim

import (
	"net/url"
	"strconv"
)

type (
	Pagination struct {
		Count      int
		StartIndex int
	}
)

func setPageToQuery(page *Pagination, query url.Values) {
	if page == nil {
		return
	}
	if page.Count != 0 {
		query.Add("count", strconv.Itoa(page.Count))
	}
	if page.StartIndex != 0 {
		query.Add("startIndex", strconv.Itoa(page.StartIndex))
	}
}
