package browse

import (
    "go-deviantart/util"
)

type MltPreview struct {
    Seed string `json:"seed"`
    Author util.User `json:"author"`
    MoreFromArtist []util.Deviation `json:"more_from_artist"`
    MoreFromDa []util.Deviation `json:"more_from_da"`
    FeaturedInCollections []gallections `json:"featured_in_collections"`
    SuggestedCollections []gallections `json:"suggested_collections"`
        
}

type gallections struct {
    Collection gallection `json:"collection"`
    Deviations []util.Deviation `json:"deviations"`
}

type gallection struct {
    Folderid int `json:"folderid"`
    Name string `json:"name"`
    Owner util.User `json:"owner"`
}

