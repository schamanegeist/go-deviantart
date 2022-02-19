package util

type Deviation struct
{
    Deviationid string `json:"deviationid"`
    Printid string `json:"printid"`
    Url string `json:"url"`
    Title string `json:"title"`
    Category string `json:"category"`
    CategoryPath string `json:"category_path"`
    Favourited bool `json:"is_favourited"`
    Deleted bool `json:"is_deleted"`
    Published bool `json:"is_published"`
    Author User `json:"author"`
    Stats map[string]int `json:"stats"`
    PublishedTime string `json:"published_time"`
    AllowsComments bool `json:"allows_comments"`
    Tier tier `json:"tier"`
    Preview preview `json:"preview"`
    Content content `json:"content"`
    Thumbs []thumb `json:"thumbs"`
    Videos []video `json:"videos"`
    Flash flash `json:"flash"`
    Daily dailyDeviation `json:"daily_deviation"`
    PremiumFolderData premiumFolder `json:"premium_folder_data"`
    TextContent editorText `json:"text_content"`
    IsPinned bool `json:"is_pinned"`
    CoverImage Deviation `json:"cover_image"`
    TierAccess interface{} `json:"tier_access"`
    PrimaryTier Deviation `json:"primary_tier"`
    Excerpt string `json:"excerpt"`
    IsMature bool `json:"is_mature"`
    IsDownloadable bool `json:"is_downloadable"`
    DownloadFilesize int `json:"download_filesize"`
    MotionBook motionBook `json:"motion_book"`
    SuggestedReasons []interface{} `json:"suggested_reasons"`
}

type preview struct {
    Src string `json:"src"`
    Height int `json:"height"`
    Width int `json:"width"`
    Transparency bool `json:"transparency"`
}

type content struct {
    Src string `json:"src"`
    Height int `json:"height"`
    Width int `json:"width"`
    Transparency bool `json:"transparency"`
    Filesize int `json:"filesize"`
}

type thumb struct {
    Src string `json:"src"`
    Height int `json:"height"`
    Width int `json:"width"`
    Transparency bool `json:"transparency"`
}

type video struct {
    Src string `json:"src"`
    Quality string `json:"quality"`
    Filesize int `json:"filesize"`
    Duration int `json:"duration"`
}

type flash struct {
    Src string `json:"src"`
    Height int `json:"height"`
    Width int `json:"width"`
}

type dailyDeviation struct {
    Body string `json:"body"`
    Time string `json:"time"`
    Giver User `json:"giver"`
    Suggester User `json:"suggester"`
}

type motionBook struct {
    EmbedUrl string `json:"embed_url"`
}

type tier struct {
    State interface{}
    IsUserSubscribed bool `json:"is_user_subscribed"`
    CanUserSubscribe bool `json:"can_user_subscribe"`
}

type premiumFolder struct {
    Type string`json:"type"`
    HasAccess bool`json:"has_access"`
    GalleryId string`json:"gallery_id"`
    PointsPrice int `json:"points_price"`
    DollarPrice float64 `json:"dollar_price"`
    NumSubscribers int `json:"num_subscribers"`
    Subproductid int `json:"subproductid"`
}

type editorText struct {
    Excerpt string `json:"excerpt"`
    Body body `json:"body"`
}

type body struct {
    Type string `json:"type"`
    Markup string `json:"markup"`
    Features string `json:"features"`
}
