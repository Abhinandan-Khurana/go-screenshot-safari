# go-screenshot-safari

## Description

This project provides a straightforward implementation of Selenium (in GoLang) with the Safari browser for capturing screenshots of web pages. It addresses the limitations posed by certain websites that require specific browsers, such as Chrome (version 96.0 or higher), Edge, or Safari, to function properly.

### Installation

```bash
go install -v github.com/Abhinandan-Khurana/go-screenshot-safari@latest
```

### Pre-requisites

Goto Safari > Develop > Developer Settings > Allow remote Automation (check this option)

## Usage

```bash

-interval_wait_time int
Time to wait between taking screenshots of URLs (in seconds) (default 1)
-load_wait_time int
Time to wait for a URL to load before taking a screenshot (in seconds) (default 2)
-output_dir string
The output directory to save screenshots (default "screenshots")
-urls_file string
Path to the file containing the list of URLs (default "urls.txt")
```

### Example usage

```bash
go run . -urls_file urls.txt -output_dir screenshots -load_wait_time 2 -interval_wait_time 1
```
