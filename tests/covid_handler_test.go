package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"line-man.com/tin-dpj/pkg/usecase" // Import the usecase package to use the mock function

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCovidSummaryHandler(t *testing.T) {
	// Create a new Gin router and set the mode to TestMode to disable Gin's panic recovery middleware.
	router := gin.New()
	router.Use(gin.TestMode)

	// Define a mock function for CountCasesByProvinceAndAge from the usecase package.
	mockCountCases := func(data []byte) (map[string]int, map[string]int, error) {
		// You can simulate the behavior of the usecase function here for testing purposes.
		// For example, return a predefined result for a given input data.
		return map[string]int{
			"Bangkok":  100,
			"Phrae":    50,
			"N/A":       5,
		}, map[string]int{
			"0-30":  30,
			"31-60": 40,
			"61+":   25,
			"N/A":    5,
		}, nil
	}

	// Replace the actual usecase function with the mock function for testing.
	usecase.CountCasesByProvinceAndAge = mockCountCases

	// Define the test cases.
	testCases := []struct {
		name           string
		requestPayload interface{} // JSON request payload
		expectedStatus int         // Expected HTTP status code
		expectedResult interface{} // Expected JSON response
	}{
		{
			name:           "Valid Request",
			requestPayload: []byte(`{"Data":[{"Age":30,"Province":"Bangkok"},{"Age":50,"Province":"Phrae"}]}`),
			expectedStatus: http.StatusOK,
			expectedResult: map[string]interface{}{
				"Province": map[string]interface{}{
					"Bangkok": 100,
					"Phrae":   50,
					"N/A":     5,
				},
				"AgeGroup": map[string]interface{}{
					"0-30":  30,
					"31-60": 40,
					"61+":   25,
					"N/A":    5,
				},
			},
		},
		// Add more test cases here.
	}

	// Run the test cases.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create an HTTP request with the given payload.
			req := httptest.NewRequest(http.MethodGet, "/covid/summary", bytes.NewBuffer(tc.requestPayload))
			req.Header.Set("Content-Type", "application/json")

			// Create an HTTP response recorder to capture the response.
			w := httptest.NewRecorder()

			// Serve the request using the router.
			router.ServeHTTP(w, req)

			// Check the response status code.
			assert.Equal(t, tc.expectedStatus, w.Code)

			// Parse the response body and compare it to the expected result.
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, response)
		})
	}
}
