package util

type User struct {
    Userid string `json:"userid"`
    Username string `json:"username"`
    Usericon string `json:"usericon"`
    Type string `json:"type"`
    IsWatching bool `json:"is_watching"`
    IsSubscribed bool `json:"is_subscribed"`
    Details details `json:"details"`
    Geo geo `json:"geo"`
    Profile profile `json:"profile"`
    Stats map[string]int `json:"stats"`
    Sidebar sidebar `json:"sidebar"`
    Session ApiSession `json:"session"`
}

type details struct {
    Sex string `json:"sex"`
    Age int `json:"age"`
    JoinDate string `json:"joindate"`
}

type geo struct {
    Country string `json:"country"`
    CountryId int `json:"countryid"`
    Timezone string `json:"timezone"`
}

type profile struct {
    UserIsArtist bool `json:"user_is_artist"`
    ArtistLevel string `json:"artist_level"`
    ArtistSpeciality string `json:"artist_speciality"`
    RealName string `json:"real_name"`
    Tagline string `json:"tagline"`
    Website string `json:"website"`
    CoverPhoto string `json:"cover_photo"`
}

type sidebar struct {
    Watched watched `json:"watched"`
}

type watched struct {
    HasNewContent bool `json:"has_new_content"`
    LinkSubnav linkSubnav `json:"link_subnav"`
    Pinned bool `json:"is_pinned"`
}

type linkSubnav struct {
    ContentType string `json:"content_type"`
}

type ApiSession struct {
    User user `json:"user"`
    Counts counts `json:"counts"`
}

type user struct {
    Userid string `json:"userid"`
    Username string `json:"username"`
    Usericon string `json:"usericon"`
    SymbolClass string `json:"symbol_class"`
}

type counts struct {
    Feedback int `json:"feedback"`
    Notes int `json:"notes"`
}
