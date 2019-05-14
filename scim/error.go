package scim

type (
	Error struct {
		Description string `json:"description"`
		Code        int    `json:"code"`
	}
)

func (e *Error) Error() string {
	return e.Description
}
