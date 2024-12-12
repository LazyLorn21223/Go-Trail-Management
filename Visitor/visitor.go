package visitor

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Visitor struct {
	Name         string
	VisitDate    string
	Trail        string
	Feedback     string
	Satisfaction string
}

var Visitors []Visitor

// Load visitor data from a CSV file
func LoadVisitorData(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
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
		visitor := Visitor{
			Name:         record[0],
			VisitDate:    record[1],
			Trail:        record[2],
			Feedback:     record[3],
			Satisfaction: record[4],
		}
		Visitors = append(Visitors, visitor)
	}
}

// Save visitor data to a CSV file
func SaveVisitorData(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, visitor := range Visitors {
		record := []string{visitor.Name, visitor.VisitDate, visitor.Trail, visitor.Feedback, visitor.Satisfaction}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}
	}
}

// Validate the satisfaction score (1-5)
func isValidSatisfaction(satisfaction string) bool {
	satisfaction = strings.TrimSpace(satisfaction)
	for _, r := range satisfaction {
		if !('0' <= r && r <= '9') {
			return false
		}
	}
	return satisfaction == "1" || satisfaction == "2" || satisfaction == "3" || satisfaction == "4" || satisfaction == "5"
}

// Validate the visit date (YYYY-MM-DD)
func isValidDate(date string) bool {
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return re.MatchString(date)
}

// Helper function to read and validate input
func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Visitor menu for managing visitors
func VisitorMenu() {
	for {
		fmt.Println("\nVisitor Tracking")
		fmt.Println("1. Add Visitor")
		fmt.Println("2. Update Visitor")
		fmt.Println("3. Delete Visitor")
		fmt.Println("4. View Visitors")
		fmt.Println("5. Back to Main Menu")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addVisitor()
		case 2:
			updateVisitor()
		case 3:
			deleteVisitor()
		case 4:
			viewVisitors()
		case 5:
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

// Add a new visitor
func addVisitor() {
	var visitor Visitor

	// Get and validate visitor name
	visitor.Name = readInput("Enter visitor name: ")
	if visitor.Name == "" {
		fmt.Println("Visitor name cannot be empty.")
		return
	}

	// Get and validate visit date
	visitor.VisitDate = readInput("Enter visit date (YYYY-MM-DD): ")
	if !isValidDate(visitor.VisitDate) {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		return
	}

	// Get and validate trail name
	visitor.Trail = readInput("Enter trail name: ")
	if visitor.Trail == "" {
		fmt.Println("Trail name cannot be empty.")
		return
	}

	// Get and validate satisfaction score
	visitor.Satisfaction = readInput("Enter satisfaction score (1-5): ")
	if !isValidSatisfaction(visitor.Satisfaction) {
		fmt.Println("Satisfaction score must be between 1 and 5.")
		return
	}

	// Get feedback
	visitor.Feedback = readInput("Enter feedback (e.g., 'satisfied, wildlife'): ")

	Visitors = append(Visitors, visitor)
	fmt.Println("Visitor added successfully.")
}

// Update an existing visitor
func updateVisitor() {
	name := readInput("Enter visitor name to update: ")
	date := readInput("Enter visit date to update: ")

	for i, visitor := range Visitors {
		if visitor.Name == name && visitor.VisitDate == date {
			// Get new visitor details with validation
			visitor.Name = readInput("Enter new name: ")
			if visitor.Name == "" {
				fmt.Println("Visitor name cannot be empty.")
				return
			}

			visitor.VisitDate = readInput("Enter new visit date (YYYY-MM-DD): ")
			if !isValidDate(visitor.VisitDate) {
				fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
				return
			}

			visitor.Trail = readInput("Enter new trail name: ")
			if visitor.Trail == "" {
				fmt.Println("Trail name cannot be empty.")
				return
			}

			visitor.Satisfaction = readInput("Enter new satisfaction score: ")
			if !isValidSatisfaction(visitor.Satisfaction) {
				fmt.Println("Satisfaction score must be between 1 and 5.")
				return
			}

			visitor.Feedback = readInput("Enter new feedback: ")

			Visitors[i] = visitor
			fmt.Println("Visitor record updated successfully.")
			return
		}
	}
	fmt.Println("Visitor not found.")
}

// Delete an existing visitor
func deleteVisitor() {
	name := readInput("Enter visitor name to delete: ")
	date := readInput("Enter visit date to delete: ")

	// Confirm before deletion
	fmt.Printf("Are you sure you want to delete the record for '%s' on '%s'? (y/n): ", name, date)
	var confirmation string
	fmt.Scanln(&confirmation)
	if confirmation != "y" {
		fmt.Println("Delete operation cancelled.")
		return
	}

	for i, visitor := range Visitors {
		if visitor.Name == name && visitor.VisitDate == date {
			Visitors = append(Visitors[:i], Visitors[i+1:]...)
			fmt.Println("Visitor record deleted successfully.")
			return
		}
	}
	fmt.Println("Visitor not found.")
}

// View all visitors
func viewVisitors() {
	if len(Visitors) == 0 {
		fmt.Println("No visitors to display.")
		return
	}

	fmt.Println("List of Visitors:")
	for _, visitor := range Visitors {
		fmt.Printf("Name: %s, Date: %s, Trail: %s, Satisfaction: %s, Feedback: %s\n", visitor.Name, visitor.VisitDate, visitor.Trail, visitor.Satisfaction, visitor.Feedback)
	}
}
