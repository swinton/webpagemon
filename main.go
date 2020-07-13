package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/swinton/webpagemon/cache"
	"github.com/swinton/webpagemon/notify"
	"github.com/swinton/webpagemon/webpage"
)

func main() {
	// Get arguments, url, selector
	url := os.Args[1]
	selector := os.Args[2]
	recipients := strings.Split(os.Args[3], ",")

	// Initialize cache
	err := cache.Init(".webpagemon")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing cache: %v\n", err)
		os.Exit(1)
	}

	// Look up previous value from cache
	cacheKey := fmt.Sprintf("%s;%s", url, selector)
	prev, err := cache.Get(cacheKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from cache: %v\n", err)
		os.Exit(1)
	}

	// Look up current value from webpage
	current, err := webpage.Get(url, selector)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", url, err)
		os.Exit(1)
	}

	// Save current value to cache
	err = cache.Set(cacheKey, current)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning writing to cache: %v\n", err)
	}

	// Compare previous with current
	if prev == current {
		fmt.Printf("Current value, same as previous: %s\n", current)
	} else {
		fmt.Printf("Current value, changed from previous: %s\n", current)

		// Notify recipients
		for _, recipient := range recipients {
			notify.Notify(recipient, fmt.Sprintf("%s: %s", url, current))
		}
	}
}
