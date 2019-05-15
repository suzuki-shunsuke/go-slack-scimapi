package scim

import (
	"encoding/json"
	"net/http"
)

func IsErrorDefault(resp *http.Response) bool {
	return resp.StatusCode >= 400
}

func ParseRespDefault(resp *http.Response, output interface{}) error {
	if output == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(output)
}

func ParseErrorRespDefault(resp *http.Response) error {
	a := &struct {
		Errors *Error
	}{}
	if err := json.NewDecoder(resp.Body).Decode(a); err != nil {
		return err
	}
	return a.Errors
}

func NewHTTPClientDefault() (*http.Client, error) {
	return &http.Client{}, nil
}
