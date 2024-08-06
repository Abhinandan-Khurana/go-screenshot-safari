package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

const (
	port = 4444
)

// sanitizeFilename replaces invalid characters in a string to create a valid filename
func sanitizeFilename(url string) string {
	replacer := strings.NewReplacer(
		":", "_",
		"/", "_",
		"\\", "_",
		"?", "_",
		"&", "_",
		"=", "_",
		"#", "_",
		"%", "_",
		" ", "_",
	)
	return replacer.Replace(url)
}

func readURLs(filePath string) ([]string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	var urls []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			urls = append(urls, trimmed)
		}
	}
	return urls, nil
}

func captureScreenshot(wd selenium.WebDriver, url string, outputFile string, loadWaitTime time.Duration) error {
	log.Printf("Navigating to URL: %s", url)
	// Navigate to the URL
	if err := wd.Get(url); err != nil {
		return fmt.Errorf("failed to load page: %v", err)
	}

	// Wait for the page to load
	log.Printf("Waiting for %v seconds for the page to load", loadWaitTime.Seconds())
	time.Sleep(loadWaitTime)

	log.Printf("Capturing screenshot for URL: %s", url)
	// Capture the screenshot
	screenshot, err := wd.Screenshot()
	if err != nil {
		return fmt.Errorf("failed to capture screenshot: %v", err)
	}

	log.Printf("Saving screenshot to file: %s", outputFile)
	// Save the screenshot to a file
	if err := ioutil.WriteFile(outputFile, screenshot, 0644); err != nil {
		return fmt.Errorf("failed to save screenshot: %v", err)
	}

	fmt.Printf("Screenshot saved to %s\n", outputFile)
	return nil
}

func main() {
	urlsFile := flag.String("urls_file", "urls.txt", "Path to the file containing the list of URLs")
	outputDir := flag.String("output_dir", "screenshots", "The output directory to save screenshots")
	loadWaitTime := flag.Int("load_wait_time", 2, "Time to wait for a URL to load before taking a screenshot (in seconds)")
	intervalWaitTime := flag.Int("interval_wait_time", 1, "Time to wait between taking screenshots of URLs (in seconds)")
	flag.Parse()

	log.Printf("Reading URLs from file: %s", *urlsFile)
	urls, err := readURLs(*urlsFile)
	if err != nil {
		log.Fatalf("Failed to read URLs: %v", err)
	}

	if _, err := os.Stat(*outputDir); os.IsNotExist(err) {
		log.Printf("Creating output directory: %s", *outputDir)
		os.Mkdir(*outputDir, os.ModePerm)
	}

	// Connect to the WebDriver instance running locally
	caps := selenium.Capabilities{"browserName": "safari"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Failed to open session: %v", err)
	}
	defer wd.Quit()

	// Maximize the browser window
	err = wd.MaximizeWindow("")
	if err != nil {
		log.Fatalf("Failed to maximize window: %v", err)
	}

	for i, url := range urls {
		sanitizedURL := sanitizeFilename(url)
		outputFile := fmt.Sprintf("%s/screenshot_%d_%s.png", *outputDir, i+1, sanitizedURL)
		log.Printf("Processing URL %d/%d: %s", i+1, len(urls), url)
		if err := captureScreenshot(wd, url, outputFile, time.Duration(*loadWaitTime)*time.Second); err != nil {
			log.Printf("Failed to capture screenshot for %s: %v", url, err)
		}
		log.Printf("Waiting for %d seconds before capturing the next screenshot", *intervalWaitTime)
		time.Sleep(time.Duration(*intervalWaitTime) * time.Second)
	}
}
