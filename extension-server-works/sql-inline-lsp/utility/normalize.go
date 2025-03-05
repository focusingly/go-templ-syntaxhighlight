package utility

import (
	"fmt"
	"net/url"
	"path/filepath"
	"runtime"
	"strings"
)

// ConvertFileURLToPath converts a file:// URL to a local file system path
func ConvertFileURLToPath(fileURL string) (string, error) {
	// Parse the URL
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return "", err
	}

	// Check if it's a file URL
	if parsedURL.Scheme != "file" {
		return "", fmt.Errorf("not a file URL: %s", fileURL)
	}

	// Decode the path
	path, err := url.PathUnescape(parsedURL.Path)
	if err != nil {
		return "", err
	}

	// Handle Windows paths
	if runtime.GOOS == "windows" {
		// Remove leading slash for Windows paths
		path = strings.TrimPrefix(path, "/")

		// Replace URL-encoded colon
		path = strings.ReplaceAll(path, "%3A", ":")
	} else {
		// For Unix-like systems, use the standard path
		path = parsedURL.Path
	}

	// Clean and normalize the path
	cleanPath := filepath.Clean(path)

	return cleanPath, nil
}
