package Feedback

import (
	"fmt"
	Visitor "project/Visitor"
	"strconv"
)

// ViewFeedbackSummary aggregates and analyzes visitor satisfaction scores.
func ViewFeedbackSummary() {
	if len(Visitor.Visitors) == 0 {
		fmt.Println("No visitor feedback available to analyze.")
		return
	}

	fmt.Println("\nFeedback Summary")

	// Satisfaction score counts and average calculation
	satisfactionCounts := make(map[int]int)
	var totalSatisfaction, totalEntries int

	for _, visitor := range Visitor.Visitors {
		score, err := strconv.Atoi(visitor.Satisfaction)
		if err != nil || score < 1 || score > 5 {
			fmt.Printf("Skipping invalid satisfaction score for visitor %s: %s\n", visitor.Name, visitor.Satisfaction)
			continue
		}

		satisfactionCounts[score]++
		totalSatisfaction += score
		totalEntries++
	}

	if totalEntries == 0 {
		fmt.Println("No valid satisfaction data to display.")
		return
	}

	// Display satisfaction counts
	fmt.Println("\nSatisfaction Score Distribution:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Score %d: %d entries\n", i, satisfactionCounts[i])
	}

	// Calculate and display average satisfaction score
	averageSatisfaction := float64(totalSatisfaction) / float64(totalEntries)
	fmt.Printf("\nAverage Satisfaction Score: %.2f\n", averageSatisfaction)
}
