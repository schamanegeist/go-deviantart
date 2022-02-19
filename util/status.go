package util

type Status struct {
    Statusid string `json:"statusid"`
    Body string `json:"body"`
    Ts string `json:"ts"`
    Url string `json:"url"`
    Comments_count integer `json:"comments_count"`
    IsShare bool `json:"is_share"`
    IsDeleted bool `json:"is_deleted"`
    Author User `json:"author"`
    Items []item `json:"items"`
    TextContent editorText `json:"text_content"`
}

type item struct {
    Type string `json:"type"`
    Status Status `json:"status"`
    Deviation Deviation `json:"deviation"`
}

