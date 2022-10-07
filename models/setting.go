package models

type Setting struct {
	Id              int    `form:"id" json:"id,omitempty"`
	SiteTitle       string `form:"site_title" json:"siteTitle,omitempty"`
	SiteLogo        string `json:"siteLogo,omitempty"`
	SiteKeywords    string `form:"site_keywords" json:"siteKeywords,omitempty"`
	SiteDescription string `form:"site_description" json:"siteDescription,omitempty"`
	NoPicture       string `json:"noPicture,omitempty"`
	SiteIcp         string `form:"site_icp" json:"siteIcp,omitempty"`
	SiteTel         string `form:"site_tel" json:"siteTel,omitempty"`
	SearchKeywords  string `form:"search_keywords" json:"searchKeywords,omitempty"`
	TongjiCode      string `form:"tongji_code" json:"tongjiCode,omitempty"`
	Appid           string `form:"appid" json:"appid,omitempty"`
	AppSecret       string `form:"app_secret" json:"appSecret,omitempty"`
	EndPoint        string `form:"end_point" json:"endPoint,omitempty"`
	BucketName      string `form:"bucket_name" json:"bucketName,omitempty"`
	OssStatus       int    `form:"oss_status" json:"ossStatus,omitempty"`
	ThumbnailSize   string `form:"thumbnail_size" json:"thumbnailSize,omitempty"`
}

func (Setting) TableName() string {
	return "setting"
}
