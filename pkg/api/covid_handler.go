package api

import (
	"io"
	"net/http"

	"line-man.com/tin-dpj/pkg/usecase"

	"github.com/gin-gonic/gin"
)

func CovidSummaryHandler(c *gin.Context) {
	// Define the URL where the JSON data is hosted.
	url := "https://static.wongnai.com/devinterview/covid-cases.json"

	// Make an HTTP GET request to fetch the JSON data.
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from the URL"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from the URL"})
		return
	}

	// Read the response body and store it in a byte slice.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Call the use case function to process the data.
	provinceCounts, ageGroupCounts, err := usecase.CountCasesByProvinceAndAge(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
		return
	}

	// Return the desired response as JSON.
	response := gin.H{
		"Province": provinceCounts,
		"AgeGroup": ageGroupCounts,
	}

	c.JSON(http.StatusOK, response)
}
