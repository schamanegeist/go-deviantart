package util

// Response struct from https://www.deviantart.com/api/v1/oauth2/placebo 
// indicating if the access_token is valid.
type Placebo struct {
    Status string `json:"status"`
}
