package external

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"signage-system/models"
	"strings"
)

var envBearerToken = os.Getenv("TWITTER_BEARER_TOKEN")

// TODO Fix : Twitter APIs are throwing 403. Huh!

// fetchWOEID fetches the Where On Earth IDentifier (WOEID) for a given location name.
// This WOEID is used to identify the location for which to fetch trending topics on Twitter.
func fetchWOEID(token, locationName string) (models.WOEIDLocation, error) {
	locations := []models.WOEIDLocation{}
	location := models.WOEIDLocation{}

	if token == "" {
		return location, fmt.Errorf("bearer token is empty")
	}

	if locationName == "" {
		return location, fmt.Errorf("location name is empty")
	}

	req, err := http.NewRequest("GET", "https://api.twitter.com/1.1/trends/available.json", nil)
	if err != nil {
		return location, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return location, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return location, fmt.Errorf("Twitter API returned status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		return location, err
	}

	if len(locations) == 0 {
		return location, fmt.Errorf("no locations found in response")
	}

	// Find the WOEID for the specified location
	for _, loc := range locations {
		if strings.EqualFold(loc.Name, locationName) {
			location = loc
			break
		}
	}

	if location.WOEID == 0 {
		return location, fmt.Errorf("location not found: %s", locationName)
	}

	return location, nil
}

// FetchTrendingTopics fetches the trending topics for a given location name.
// It uses the bearer token for authentication and the WOEID to identify the location.
func FetchTrendingTopics(location string) ([]models.TrendingTopic, error) {
	bearerToken := envBearerToken
	if bearerToken == "" {
		return nil, fmt.Errorf("bearer token is empty")
	}

	if location == "" {
		return nil, fmt.Errorf("location is empty")
	}

	loc, err := fetchWOEID(bearerToken, location)
	if err != nil {
		log.Printf("Error while fetching WOEID: %v", err)
		return nil, err
	}

	// Fetch trending topics for the location identified by the WOEID
	url := fmt.Sprintf("https://api.twitter.com/1.1/trends/place.json?id=%v", loc.WOEID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Twitter API returned status: %s", resp.Status)
	}

	var trendsResponse []models.TrendsResponse
	if err := json.NewDecoder(resp.Body).Decode(&trendsResponse); err != nil {
		return nil, err
	}

	if len(trendsResponse) == 0 {
		return nil, fmt.Errorf("no trending topics found")
	}

	if len(trendsResponse[0].Trends) == 0 {
		return nil, fmt.Errorf("no trends found for location: %s", location)
	}

	return trendsResponse[0].Trends, nil
}

// FetchTweetsForTopic fetches popular tweets for a given topic.
// It uses the bearer token for authentication and searches for tweets matching the topic.
func FetchTweetsForTopic(topic string) ([]models.Tweet, error) {
	bearerToken := envBearerToken
	if bearerToken == "" {
		return nil, fmt.Errorf("bearer token is empty")
	}

	if topic == "" {
		return nil, fmt.Errorf("topic is empty")
	}

	// Search for popular tweets containing the topic
	url := fmt.Sprintf("https://api.twitter.com/1.1/search/tweets.json?q=%s&result_type=popular&count=10", topic)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Twitter API returned status: %s", resp.Status)
	}

	var tweetsResponse models.TweetsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tweetsResponse); err != nil {
		return nil, err
	}

	if len(tweetsResponse.Statuses) == 0 {
		return nil, fmt.Errorf("no tweets found for topic: %s", topic)
	}

	return tweetsResponse.Statuses, nil
}
