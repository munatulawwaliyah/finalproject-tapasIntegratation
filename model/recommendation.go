package model

import (
	"strings"
)

func GenerateRecommendation(response Response) string {
    // Example recommendation logic
    if response.Aggregator == "total" {
        return "Based on the total provided, consider optimizing your resource allocation."
    }
    if strings.Contains(response.Answer, "increase") {
        return "It seems there is an increase. You might want to investigate the cause and take appropriate actions."
    }
    return "No specific recommendation available."
}
