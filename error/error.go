package error

import (
    "fmt"
)

type ApiError struct {
    Type string `json:"error"`
    Desc string `json:"error_description"`
    Code int `json:"error_code"`
    Details map[string]string `json:"error_details"`
}

func (e ApiError) Error() string {
    return fmt.Sprintf("%s: %s", e.Type, e.Desc)
}
