package models

type Content struct {
	ID          string `json:"id"`
	CreatedTime int64  `json:"created_time"`
	Location    string `json:"location"`
	Weather     string `json:"weather"`
	SocialMedia string `json:"social_media"`
}
