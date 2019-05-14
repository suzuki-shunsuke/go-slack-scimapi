package scim

import (
	"encoding/json"
	"net/http"
)

func isError(resp *http.Response) bool {
	return resp.StatusCode >= 400
}

func parseResp(resp *http.Response, output interface{}) error {
	if output == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(output)
}

func parseErrorResp(resp *http.Response) error {
	a := &struct {
		Errors *Error
	}{}
	if err := json.NewDecoder(resp.Body).Decode(a); err != nil {
		return err
	}
	return a.Errors
}

func clientFn() (*http.Client, error) {
	return &http.Client{}, nil
}
