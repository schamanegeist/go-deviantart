package browse

import (
    "go-deviantart/util"
)

type Topics struct {
    Has_more bool `json:"has_more"`
    Next_offset int `json:"next_offset"`
    Next_cursor string `json:"next_cursor"`
    Prev_cursor string `json:"prev_cursor"`
    Results []result `json:"results"`
    Session util.ApiSession `json:"session"`
}

type TopTopics struct {
    Results []result `json:"result"`
    Session util.ApiSession `json:"session"`
}

type result struct {
    Name string `json:"name"`
    CanonicalName string `json:"canonical_name"`
    Examples []util.Deviation `json:"example_deviations"`
    Deviations []util.Deviation `json:"deviations"`
}
