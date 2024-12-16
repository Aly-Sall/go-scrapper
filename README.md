Go Web Scraper

A simple web scraping tool built with Go. This program extracts headings and links from any website and saves the data to both CSV and JSON files.
⚠️ Important: This tool is intended for personal and legal use only. Please ensure you comply with the website's terms of service and applicable laws before scraping any data.


Features :

Extracts headings (h1, h2, h3, etc.) and links (<a> tags) from a given website.
Saves scraped data to:
CSV file (output.csv)
JSON file (output.json)
Logs scraping activity to a file (scraper.log) for debugging and tracking.
Handles errors gracefully with clear logs.
Can scrape multiple pages (pagination handling).
Uses Go routines for concurrent scraping (improved performance).
Can handle JavaScript-rendered websites using tools like chromedp.

Installation

1. Clone the repository:

2. Install dependencies:
Install Go modules: go mod tidy

3.Build the project: go build -o scraper

4. Run the scraper: ./scraper

Usage

When you run the program, it will prompt you to enter the URL of the website you want to scrape.
Example: https://example.com
The tool will extract headings and links from the website and save the results to two files:
output.csv
output.json
Logs will be written to scraper.log for tracking activity and debugging.

Legal Disclaimer

This tool is designed for educational purposes and personal use. Scraping websites may violate their terms of service or applicable laws if done without permission.
Check the website's robots.txt file to understand what is allowed for crawling and scraping.
Do not scrape sensitive, copyrighted, or private data without explicit consent.
Use this tool responsibly and at your own risk.

Technical Details

Core Libraries
PuerkitoBio/goquery: For parsing and navigating HTML.
encoding/csv: For saving data to CSV files.
encoding/json: For saving data to JSON files.
Concurrency
The scraper uses Go routines to handle multiple URLs simultaneously for faster data collection.
Dynamic Content
For JavaScript-rendered websites, consider using:
chromedp: Headless Chrome integration for Go.

Contribution

Feel free to fork this repository and suggest improvements via pull requests. Contributions are always welcome! 🙏 

License

This project is licensed under the MIT License. See the LICENSE file for details.

