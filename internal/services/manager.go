package services

import (
	"golang-console-project/internal/github"
	"golang-console-project/internal/opslevel"
	"log"
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
