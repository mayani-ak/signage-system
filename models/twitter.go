package models

type WOEIDLocation struct {
	Name  string `json:"name"`
	WOEID int    `json:"woeid"`
}

type BearerTokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

type TrendingTopic struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type TrendsResponse struct {
	Trends []TrendingTopic `json:"trends"`
}

type Tweet struct {
	ID   string `json:"id_str"`
	Text string `json:"text"`
}

type TweetsResponse struct {
	Statuses []Tweet `json:"statuses"`
}
