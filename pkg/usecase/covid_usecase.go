package usecase

import (
	"encoding/json"
)

func CountCasesByProvinceAndAge(data []byte) (map[string]int, map[string]int, error) {
    // Define the data structures for the desired response.
    provinceCounts := make(map[string]int)
    ageGroupCounts := map[string]int{
        "0-30": 0,
        "31-60": 0,
        "61+": 0,
    }

    // Parse the JSON data into a struct.
    var covidData struct {
        Data []struct {
            Age      int    `json:"Age"`
            Province string `json:"Province"`
        } `json:"Data"`
    }

    if err := json.Unmarshal(data, &covidData); err != nil {
        return nil, nil, err
    }

    // Process the data and count cases by province and age group.
    for _, entry := range covidData.Data {
        if entry.Province == "" {
            entry.Province = "N/A"
        }

        if entry.Age >= 0 && entry.Age <= 30 {
            ageGroupCounts["0-30"]++
        } else if entry.Age >= 31 && entry.Age <= 60 {
            ageGroupCounts["31-60"]++
        } else if entry.Age >= 61 {
            ageGroupCounts["61+"]++
        } else {
            // Handle cases with no age data.
            ageGroupCounts["N/A"]++
        }

        // Count cases by province.
        provinceCounts[entry.Province]++
    }

    return provinceCounts, ageGroupCounts, nil
}
