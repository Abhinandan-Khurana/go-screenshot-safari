# Endpoint_screenshot_safari_browser

## Description:

### Pre-requisites:

Goto Safari > Develop > Developer Settings > Allow remote Automation (check this option)

## Usage:

-interval_wait_time int
Time to wait between taking screenshots of URLs (in seconds) (default 1)
-load_wait_time int
Time to wait for a URL to load before taking a screenshot (in seconds) (default 2)
-output_dir string
The output directory to save screenshots (default "screenshots")
-urls_file string
Path to the file containing the list of URLs (default "urls.txt")

### Example usage:

```bash
go run . -urls_file urls.txt -output_dir screenshots -load_wait_time 2 -interval_wait_time 1
```
