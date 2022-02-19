package browse

import (
    "go-deviantart/util"
)

type BrowseResult struct {
    HasMore bool `json:"has_more"`
    NextOffset int `json:"next_offset"`
    NextCursor string `json:"next_cursor"`
    PrevCursor string `json:"prev_cursor"`
    EffectivePageType string `json:"effective_page_type"`
    ErrorCode int `json:"error_code"`
    EstimatedTotal int `json:"estimated_total"`
    Results []util.Deviation `json:"results"`
    Session util.ApiSession `json:"session"`
}
