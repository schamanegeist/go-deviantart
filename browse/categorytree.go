package browse

type CategoryTree struct {
    Categories []category `json:"categories"`
}

type category struct {
    Path string `json:"catpath"`
    Title string `json:"title"`
    HasSubcategory bool `json:"has_subcategory"`
    Parent string `json:"parent_catpath"`
}
