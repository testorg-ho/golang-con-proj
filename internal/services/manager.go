package services

import (
	"golang-console-project/internal/github"
	"golang-console-project/internal/opslevel"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func ManageServices(opsClient opslevel.OpsLevelClient, gitClient github.GitHubClient) error {
	services, err := opsClient.GetServices()
	if err != nil {
		return err
	}

	for _, service := range services {
		err := gitClient.CreateRepo(service.Name)
		if err != nil {
			log.Printf("Failed to create repository for service %s: %v", service.Name, err)
		} else {
			log.Printf("Successfully created repository for service %s", service.Name)
		}
	}
	return nil
}

// DeleteDirsWithYesterdayDate deletes directories in the baseDir with names matching the pattern {some-name}-YYYY-MM-DD
// and containing yesterday's date.
func DeleteDirsWithYesterdayDate(baseDir string) error {
	// Define the regex pattern for directory names
	pattern := regexp.MustCompile(`^[a-zA-Z0-9_-]+-\d{4}-\d{2}-\d{2}$`)

	// Calculate yesterday's date
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	// Walk through the base directory
	return filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the current item is a directory and matches the pattern
		if info.IsDir() && pattern.MatchString(info.Name()) {
			// Check if the directory name ends with yesterday's date
			if filepath.Base(info.Name()) == yesterday {
				// Delete the directory
				if err := os.RemoveAll(path); err != nil {
					return err
				}
			}
		}
		return nil
	})
}



// find /path/to/base/directory -type d -name "*-$(date -v-1d +%Y-%m-%d)" -exec rm -rf {} +
