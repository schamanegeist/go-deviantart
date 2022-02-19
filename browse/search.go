package browse

type Search struct {
    Results []tag `json:"results"`
    Session util.ApiSession `json:"session"`
}

type tag struct {
    TagName string `json:"tag_name "`
}
