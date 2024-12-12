package Maintenance

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Maintenance struct {
	TrailName string
	Date      string
	Type      string
}

var MaintenanceRecords []Maintenance

// Load maintenance data from a CSV file
func LoadMaintenanceData(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for _, record := range records {
		if len(record) < 3 {
			fmt.Println("Skipping invalid record:", record)
			continue
		}

		// Validate and parse date
		if !isValidDate(record[1]) {
			fmt.Println("Skipping record with invalid date:", record)
			continue
		}

		maintenance := Maintenance{
			TrailName: record[0],
			Date:      record[1],
			Type:      record[2],
		}
		MaintenanceRecords = append(MaintenanceRecords, maintenance)
	}
}

// Save maintenance data to a CSV file
func SaveMaintenanceData(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range MaintenanceRecords {
		data := []string{record.TrailName, record.Date, record.Type}
		if err := writer.Write(data); err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}
	}
}

// Validate if the date is in the correct format (YYYY-MM-DD)
func isValidDate(date string) bool {
	// Using regex to match the date format (YYYY-MM-DD)
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return re.MatchString(date)
}

// Maintenance menu for managing maintenance records
func MaintenanceMenu() {
	for {
		fmt.Println("\nMaintenance Scheduling")
		fmt.Println("1. Add Maintenance Record")
		fmt.Println("2. Update Maintenance Record")
		fmt.Println("3. Delete Maintenance Record")
		fmt.Println("4. View Maintenance Records")
		fmt.Println("5. Back to Main Menu")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addMaintenance()
		case 2:
			updateMaintenance()
		case 3:
			deleteMaintenance()
		case 4:
			viewMaintenanceRecords()
		case 5:
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

// Add a new maintenance record
func addMaintenance() {
	var record Maintenance
	reader := bufio.NewReader(os.Stdin)

	// Get trail name with validation
	fmt.Print("Enter trail name: ")
	record.TrailName, _ = reader.ReadString('\n')
	record.TrailName = strings.TrimSpace(record.TrailName)
	if record.TrailName == "" {
		fmt.Println("Trail name cannot be empty.")
		return
	}

	// Get date with validation
	fmt.Print("Enter maintenance date (YYYY-MM-DD): ")
	record.Date, _ = reader.ReadString('\n')
	record.Date = strings.TrimSpace(record.Date)
	if !isValidDate(record.Date) {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		return
	}

	// Get maintenance type
	fmt.Print("Enter maintenance type (e.g., cleaning, repair): ")
	record.Type, _ = reader.ReadString('\n')
	record.Type = strings.TrimSpace(record.Type)
	if record.Type == "" {
		fmt.Println("Maintenance type cannot be empty.")
		return
	}

	// Add the maintenance record
	MaintenanceRecords = append(MaintenanceRecords, record)
	fmt.Println("Maintenance record added successfully.")
}

// Update an existing maintenance record
func updateMaintenance() {
	var trailName, date string
	reader := bufio.NewReader(os.Stdin)

	// Get trail name and date to uniquely identify the record
	fmt.Print("Enter the trail name of the maintenance record to update: ")
	trailName, _ = reader.ReadString('\n')
	trailName = strings.TrimSpace(trailName)

	// Get maintenance date
	fmt.Print("Enter the maintenance date (YYYY-MM-DD): ")
	date, _ = reader.ReadString('\n')
	date = strings.TrimSpace(date)
	if !isValidDate(date) {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		return
	}

	// Find the record by trail name and date
	for i, record := range MaintenanceRecords {
		if record.TrailName == trailName && record.Date == date {
			// Update the maintenance record details
			fmt.Print("Enter new maintenance type: ")
			record.Type, _ = reader.ReadString('\n')
			record.Type = strings.TrimSpace(record.Type)
			if record.Type == "" {
				fmt.Println("Maintenance type cannot be empty.")
				return
			}

			// Update the record in the MaintenanceRecords slice
			MaintenanceRecords[i] = record
			fmt.Println("Maintenance record updated successfully.")

			// Save the updated data back to the CSV file
			SaveMaintenanceData("data/maintenance.csv")
			return
		}
	}

	// If the maintenance record wasn't found
	fmt.Println("Maintenance record not found.")
}

// Delete an existing maintenance record
func deleteMaintenance() {
	var trailName, date string
	reader := bufio.NewReader(os.Stdin)

	// Get trail name
	fmt.Print("Enter the trail name of the maintenance record to delete: ")
	trailName, _ = reader.ReadString('\n')
	trailName = strings.TrimSpace(trailName)

	// Get date
	fmt.Print("Enter the maintenance date (YYYY-MM-DD): ")
	date, _ = reader.ReadString('\n')
	date = strings.TrimSpace(date)
	if !isValidDate(date) {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		return
	}

	// Confirm before deletion
	fmt.Printf("Are you sure you want to delete the maintenance record for '%s' on '%s'? (y/n): ", trailName, date)
	var confirmation string
	fmt.Scanln(&confirmation)
	if confirmation != "y" {
		fmt.Println("Delete operation cancelled.")
		return
	}

	// Find and delete the record
	found := false
	for i, record := range MaintenanceRecords {
		if record.TrailName == trailName && record.Date == date {
			MaintenanceRecords = append(MaintenanceRecords[:i], MaintenanceRecords[i+1:]...)
			found = true
			fmt.Println("Maintenance record deleted successfully.")
			break
		}
	}

	if !found {
		fmt.Println("Maintenance record not found.")
		return
	}

	// Save the updated data to the CSV file
	SaveMaintenanceData("data/maintenance.csv")
}

// View all maintenance records
func viewMaintenanceRecords() {
	if len(MaintenanceRecords) == 0 {
		fmt.Println("No maintenance records to display.")
		return
	}

	fmt.Println("Maintenance Records:")
	for _, record := range MaintenanceRecords {
		fmt.Printf("Trail Name: %s, Date: %s, Type: %s\n",
			record.TrailName, record.Date, record.Type)
	}
}
