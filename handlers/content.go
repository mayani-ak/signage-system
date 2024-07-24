package handlers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"signage-system/external"
	"signage-system/firestore"
	"signage-system/models"
	"time"

	"github.com/labstack/echo/v4"
)

var location = "Seattle" // TODO accept location in request

// UpdateContent fetches weather data, trending topics, and tweets, then updates the Firestore with the new content.
func UpdateContent(c echo.Context) error {
	// Fetch weather data for the specified location
	weather, err := external.FetchWeather(location)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	/*
		// Fetch trending topics for the specified location
		trendingTopics, err := external.FetchTrendingTopics(location)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		// Check if there are any trending topics
		if len(trendingTopics) == 0 {
			return c.JSON(http.StatusInternalServerError, "No trending topics found")
		}

		// Fetch tweets for the most trending topic
		tweets, err := external.FetchTweetsForTopic(trendingTopics[0].Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		// Check if there are any tweets
		if len(tweets) == 0 {
			return c.JSON(http.StatusInternalServerError, "No tweets found for trending topic")
		}
	*/

	// Create content struct with fetched data
	content := models.Content{
		ID:          uuid.New().String(),
		CreatedTime: time.Now().Unix(),
		Location:    location,
		Weather:     fmt.Sprintf("%0.1f°C, %s", weather.Current.Temperature, weather.Current.Condition.Text),
		//	SocialMedia: tweets[0].Text, // Assuming we want the first tweet from the most trending topic
	}

	// Add the new content to the Firestore collection
	_, _, err = firestore.Client.Collection("content").Add(context.Background(), content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, content)
}

// GetContent retrieves all content documents from the Firestore collection.
func GetContent(c echo.Context) error {

	query := firestore.Client.Collection("content").Where("Location", "==", location).Limit(10)
	// Retrieve all documents from the given location
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Check if there are any documents
	if len(docs) == 0 {
		return c.JSON(http.StatusNotFound, "No content found")
	}

	var contents []models.Content
	for _, doc := range docs {
		var content models.Content
		// Map the Firestore document data to the content struct
		if err := doc.DataTo(&content); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		contents = append(contents, content)
	}

	return c.JSON(http.StatusOK, contents)
}
