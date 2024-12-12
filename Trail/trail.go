package Trail

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var TrailRecords []Trail

type Trail struct {
	Name       string
	Location   string
	Difficulty string
	Length     float64
	Status     string
}

// Load trail data from a CSV file
func LoadTrailData(filePath string) {
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
		if len(record) < 5 {
			fmt.Println("Skipping invalid record:", record)
			continue
		}
		length, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Println("Error parsing trail length:", err)
			continue
		}
		trail := Trail{
			Name:       record[0],
			Location:   record[1],
			Difficulty: record[2],
			Length:     length,
			Status:     record[4],
		}
		TrailRecords = append(TrailRecords, trail)
	}
}

// Save trail data to a CSV file
func SaveTrailData(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, trail := range TrailRecords {
		record := []string{trail.Name, trail.Location, trail.Difficulty, strconv.FormatFloat(trail.Length, 'f', 2, 64), trail.Status}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}
	}
}

// Trail menu for managing trails
func TrailMenu() {
	for {
		fmt.Println("\nManage Trails")
		fmt.Println("1. Add Trail")
		fmt.Println("2. Update Trail")
		fmt.Println("3. Delete Trail")
		fmt.Println("4. View Trails")
		fmt.Println("5. Back to Main Menu")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTrail()
		case 2:
			updateTrail()
		case 3:
			deleteTrail()
		case 4:
			viewTrails()
		case 5:
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

// Add a new trail
func addTrail() {
	reader := bufio.NewReader(os.Stdin)

	var trail Trail

	// Get trail name
	fmt.Print("Enter trail name: ")
	trail.Name, _ = reader.ReadString('\n')
	trail.Name = strings.TrimSpace(trail.Name)
	if trail.Name == "" {
		fmt.Println("Trail name cannot be empty.")
		return
	}

	// Get trail location
	fmt.Print("Enter location: ")
	trail.Location, _ = reader.ReadString('\n')
	trail.Location = strings.TrimSpace(trail.Location)
	if trail.Location == "" {
		fmt.Println("Location cannot be empty.")
		return
	}

	// Check for duplicates
	for _, existingTrail := range TrailRecords {
		if strings.EqualFold(existingTrail.Name, trail.Name) && strings.EqualFold(existingTrail.Location, trail.Location) {
			fmt.Printf("A trail with the name '%s' and location '%s' already exists.\n", trail.Name, trail.Location)
			return
		}
	}

	// Get difficulty
	fmt.Print("Enter difficulty (Easy, Medium, Hard): ")
	trail.Difficulty, _ = reader.ReadString('\n')
	trail.Difficulty = strings.TrimSpace(trail.Difficulty)
	if trail.Difficulty == "" {
		fmt.Println("Difficulty cannot be empty.")
		return
	}

	// Get length with validation
	fmt.Print("Enter length (miles): ")
	for {
		_, err := fmt.Scanln(&trail.Length)
		if err != nil || trail.Length <= 0 {
			fmt.Println("Please enter a valid positive number for trail length.")
			fmt.Print("Enter length (miles): ")
			continue
		}
		break
	}

	// Get status
	fmt.Print("Enter status (open/closed): ")
	trail.Status, _ = reader.ReadString('\n')
	trail.Status = strings.TrimSpace(trail.Status)
	if trail.Status == "" {
		fmt.Println("Status cannot be empty.")
		return
	}

	// Add the trail
	TrailRecords = append(TrailRecords, trail)
	fmt.Println("Trail added successfully.")
}

// Update an existing trail
func updateTrail() {
	reader := bufio.NewReader(os.Stdin)

	// Get the trail name and location to uniquely identify the trail
	fmt.Print("Enter the name of the trail to update: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter the location of the trail to update: ")
	location, _ := reader.ReadString('\n')
	location = strings.TrimSpace(location)

	// Search for the trail by name and location
	for i, trail := range TrailRecords {
		if trail.Name == name && trail.Location == location {
			// Get new details for the trail
			fmt.Print("Enter new location: ")
			trail.Location, _ = reader.ReadString('\n')
			trail.Location = strings.TrimSpace(trail.Location)

			fmt.Print("Enter new difficulty: ")
			trail.Difficulty, _ = reader.ReadString('\n')
			trail.Difficulty = strings.TrimSpace(trail.Difficulty)

			fmt.Print("Enter new length (miles): ")
			fmt.Scanln(&trail.Length)

			fmt.Print("Enter new status (open/closed): ")
			trail.Status, _ = reader.ReadString('\n')
			trail.Status = strings.TrimSpace(trail.Status)

			// Update the trail record in memory
			TrailRecords[i] = trail
			fmt.Println("Trail updated successfully.")

			// Save the updated data back to the CSV file
			SaveTrailData("data/trails.csv")
			return
		}
	}
	fmt.Println("Trail not found.")
}

// Delete an existing trail
// Delete an existing trail
func deleteTrail() {
	reader := bufio.NewReader(os.Stdin)

	// Get trail name with spaces
	fmt.Print("Enter the name of the trail to delete: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Get location with spaces
	fmt.Print("Enter the location of the trail to delete: ")
	location, _ := reader.ReadString('\n')
	location = strings.TrimSpace(location)

	// Confirm before deletion
	fmt.Printf("Are you sure you want to delete the trail '%s' at location '%s'? (y/n): ", name, location)
	var confirmation string
	fmt.Scanln(&confirmation)
	if confirmation != "y" {
		fmt.Println("Delete operation cancelled.")
		return
	}

	// Find and delete the trail by name and location
	for i, trail := range TrailRecords {
		if trail.Name == name && trail.Location == location {
			// Delete the trail record from TrailRecords
			TrailRecords = append(TrailRecords[:i], TrailRecords[i+1:]...)
			fmt.Println("Trail deleted successfully.")
			// Save the updated data back to the CSV file
			SaveTrailData("data/trails.csv")
			return
		}
	}
	fmt.Println("Trail not found.")
}

// View all trails
func viewTrails() {
	if len(TrailRecords) == 0 {
		fmt.Println("No trails to display.")
		return
	}

	fmt.Println("Trail List:")
	for _, trail := range TrailRecords {
		fmt.Printf("Name: %s, Location: %s, Difficulty: %s, Length: %.2f miles, Status: %s\n",
			trail.Name, trail.Location, trail.Difficulty, trail.Length, trail.Status)
	}
}
