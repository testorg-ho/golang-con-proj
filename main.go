package main

import (
	"golang-console-project/internal/github"
	"golang-console-project/internal/opslevel"
	"golang-console-project/internal/services"
	"log"
)

func main() {
	opsClient := &opslevel.DefaultOpsLevelClient{}
	gitClient := &github.DefaultGitHubClient{}

	err := services.ManageServices(opsClient, gitClient)
	if err != nil {
		log.Fatalf("Error managing services: %v", err)
	}

	log.Println("All services processed successfully.")
}
