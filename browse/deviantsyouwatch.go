package browse

import (
    "go-deviantart/util"
)

type DywPosts struct {
    HasMore bool `json:"has_more"`
    NextOffset int `json:"next_offset"`
    ErrorCode int `json:"error_code"`
    Results []journal `json:"results"`
    Session util.ApiSession `json:"session"`
}

type journal struct {
    Journal util.Deviation `json:"journal"`
    Status util.Status `json:"status"`
}
