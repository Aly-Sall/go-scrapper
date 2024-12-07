package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Log setup
func init() {
	// Create or open a log file
	file, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error creating log file: %v", err)
	}
	// Set output of logs to file
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Scraper started")
}

type ScrapedData struct {
	Headings []string `json:"headings"`
	Links    []string `json:"links"`
}

func scrapeWebsite(url string) (*ScrapedData, error) {
	log.Printf("Starting scrape for URL: %s\n", url)

	// Fetch the URL
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Printf("Error fetching URL %s: %v\n", url, err)
		return nil, err
	}

	data := &ScrapedData{}
	// Extract headings
	doc.Find("h1, h2, h3, h4, h5, h6").Each(func(i int, s *goquery.Selection) {
		heading := strings.TrimSpace(s.Text())
		if heading != "" {
			data.Headings = append(data.Headings, heading)
		}
	})
	log.Printf("Extracted %d headings from %s\n", len(data.Headings), url)

	// Extract links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			data.Links = append(data.Links, link)
		}
	})
	log.Printf("Extracted %d links from %s\n", len(data.Links), url)

	return data, nil
}

func saveToCSV(data *ScrapedData, filename string) error {
	log.Printf("Saving data to CSV file: %s\n", filename)

	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Error creating CSV file: %v\n", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headings
	if err := writer.Write([]string{"Headings"}); err != nil {
		log.Printf("Error writing to CSV file: %v\n", err)
		return err
	}
	for _, heading := range data.Headings {
		if err := writer.Write([]string{heading}); err != nil {
			log.Printf("Error writing heading to CSV file: %v\n", err)
		}
	}

	// Write links
	if err := writer.Write([]string{"Links"}); err != nil {
		log.Printf("Error writing to CSV file: %v\n", err)
		return err
	}
	for _, link := range data.Links {
		if err := writer.Write([]string{link}); err != nil {
			log.Printf("Error writing link to CSV file: %v\n", err)
		}
	}

	log.Printf("Data successfully saved to CSV: %s\n", filename)
	return nil
}

func saveToJSON(data *ScrapedData, filename string) error {
	log.Printf("Saving data to JSON file: %s\n", filename)

	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Error creating JSON file: %v\n", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		log.Printf("Error encoding JSON data: %v\n", err)
		return err
	}

	log.Printf("Data successfully saved to JSON: %s\n", filename)
	return nil
}

func main() {
	var url string
	fmt.Print("Enter the URL to scrape: ")
	fmt.Scanln(&url)

	data, err := scrapeWebsite(url)
	if err != nil {
		log.Fatalf("Failed to scrape website: %v\n", err)
	}

	// Save results to files
	if err := saveToCSV(data, "output.csv"); err != nil {
		log.Printf("Error saving CSV: %v\n", err)
	}
	if err := saveToJSON(data, "output.json"); err != nil {
		log.Printf("Error saving JSON: %v\n", err)
	}

	log.Println("Scraping completed successfully")
	fmt.Println("Scraping completed. Check output.csv, output.json, and scraper.log for results.")
}
