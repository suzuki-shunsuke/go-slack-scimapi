package scim

type (
	// Error is Slack SCIM API's error response body.
	// https://api.slack.com/scim#errors
	Error struct {
		Description string `json:"description"`
		Code        int    `json:"code"`
	}
)

// Error returns an error's description.
func (e *Error) Error() string {
	return e.Description
}
