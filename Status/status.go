package Status

import (
	"fmt"
	"project/Maintenance"
	"project/Trail"
	"time"
)

// Load the data once at the beginning of the program or before displaying the status
func LoadData() {
	// Load trail and maintenance data before any interaction
	Trail.LoadTrailData("data/trails.csv")
	Maintenance.LoadMaintenanceData("data/maintenance.csv")
}

// ViewTrailStatus displays the status and maintenance information of all trails
func ViewTrailStatus() {
	// Check if the data has been loaded
	if len(Trail.TrailRecords) == 0 {
		fmt.Println("No trail data available.")
		return
	}

	if len(Maintenance.MaintenanceRecords) == 0 {
		fmt.Println("No maintenance data available.")
		return
	}

	// Print trail status summary
	fmt.Println("\nTrail Status Summary:")

	// Loop through trails and display their status and maintenance info
	for _, trail := range Trail.TrailRecords {
		latestMaintenance, found := getLastMaintenance(trail.Name)
		// Display trail info only once
		fmt.Printf("Trail Name: %s\n", trail.Name)
		fmt.Printf("Location: %s\n", trail.Location)
		fmt.Printf("Status: %s\n", trail.Status)

		// Display maintenance info if found
		if found {
			fmt.Printf("Last Maintained: %s\n", latestMaintenance.Date)
			fmt.Printf("Maintenance Type: %s\n", latestMaintenance.Type)
		} else {
			fmt.Println("Maintenance Record: No Maintenance Found")
		}
		fmt.Println() // Adds a blank line between each trail's information
	}
}

func getLastMaintenance(trailName string) (Maintenance.Maintenance, bool) {
	var latest Maintenance.Maintenance
	found := false

	// Iterate over maintenance records to find the most recent maintenance record for the trail
	for _, record := range Maintenance.MaintenanceRecords {
		if record.TrailName == trailName {
			recordDate, _ := time.Parse("2006-01-02", record.Date)
			latestDate, _ := time.Parse("2006-01-02", latest.Date)

			if !found || recordDate.After(latestDate) {
				latest = record
				found = true
			}
		}
	}
	return latest, found
}
