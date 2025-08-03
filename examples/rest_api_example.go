package main

import (
	"brokolisql-go/pkg/common"
	"brokolisql-go/pkg/fetchers"
	"fmt"
	"log"
	"os"
)

// This example demonstrates how to use the REST API fetcher to retrieve data
// from a public API and process it.
func main() {
	// Create a REST fetcher
	fetcherType := "rest"
	fetcher, err := fetchers.GetFetcher(fetcherType)
	if err != nil {
		log.Fatalf("Failed to create fetcher: %v", err)
	}

	// Define the API endpoint to fetch data from
	// This is a public API that returns a list of users
	apiURL := "https://jsonplaceholder.typicode.com/users"

	// Optional: Define additional options for the request
	options := map[string]interface{}{
		"method": "GET",
		"headers": map[string]string{
			"Accept": "application/json",
		},
	}

	// Fetch data from the API
	fmt.Println("Fetching data from:", apiURL)
	dataset, err := fetcher.Fetch(apiURL, options)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}

	// Print the columns and number of rows
	fmt.Printf("Fetched %d rows with columns: %v\n", len(dataset.Rows), dataset.Columns)

	// Print the first row as an example
	if len(dataset.Rows) > 0 {
		fmt.Println("\nFirst row data:")
		for key, value := range dataset.Rows[0] {
			fmt.Printf("  %s: %v\n", key, value)
		}
	}

	// Example of how to save the fetched data to a file
	// This could be used to then load the data using the existing loaders
	fmt.Println("\nSaving data to users.json...")
	saveToFile(dataset, "users.json")
	fmt.Println("Data saved. You can now use the JSON loader to process this file.")
}

// Helper function to save the dataset to a JSON file
func saveToFile(dataset *common.DataSet, filename string) error {
	// Create a JSON representation of the data
	// For simplicity, we'll just create a basic JSON array
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write opening bracket for JSON array
	file.WriteString("[\n")

	// Write each row as a JSON object
	for i, row := range dataset.Rows {
		file.WriteString("  {")

		// Write each field
		j := 0
		for key, value := range row {
			file.WriteString(fmt.Sprintf("\n    %q: %q", key, fmt.Sprintf("%v", value)))
			if j < len(row)-1 {
				file.WriteString(",")
			}
			j++
		}

		file.WriteString("\n  }")

		// Add comma if not the last row
		if i < len(dataset.Rows)-1 {
			file.WriteString(",")
		}
		file.WriteString("\n")
	}

	// Write closing bracket for JSON array
	file.WriteString("]\n")

	return nil
}
