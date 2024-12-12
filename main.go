package main

// By Lauren Auer

import (
	"fmt"
	"os"
	Feedback "project/Feedback"
	Maintenance "project/Maintenance"
	Status "project/Status"
	Trail "project/Trail"
	Visitor "project/Visitor"
)

func main() {
	// Load all necessary data files once at the start
	Status.LoadData() // This will load both Trail and Maintenance data

	// Main menu loop
	for {
		fmt.Println("\nWelcome to the Trail Management System")
		fmt.Println("\nSelect an option below")
		fmt.Println("1. Manage Trails")
		fmt.Println("2. Track Visitors")
		fmt.Println("3. Track Maintenance")
		fmt.Println("4. Feedback Summary")
		fmt.Println("5. Trail Status")
		fmt.Println("6. Save and Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			Trail.TrailMenu()
		case 2:
			Visitor.VisitorMenu()
		case 3:
			Maintenance.MaintenanceMenu()
		case 4:
			Feedback.ViewFeedbackSummary()
		case 5:
			Status.ViewTrailStatus()
		case 6:
			saveAndExit()
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func saveAndExit() {
	// Save all data before exiting
	Trail.SaveTrailData("data/trails.csv")
	Visitor.SaveVisitorData("data/visitors.csv")
	Maintenance.SaveMaintenanceData("data/maintenance.csv")
	fmt.Println("Data saved. Exiting application.")
	os.Exit(0)
}
