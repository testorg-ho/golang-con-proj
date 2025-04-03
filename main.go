package main

import (
	"golang-console-project/internal/github"
	"golang-console-project/internal/opslevel"
	"golang-console-project/internal/services"
	"log"
)

func main() {

	if say { // Use the constant from the same package
		log.Println(greet) // Use the constant from the same package
	}
	// Initialize OpsLevel and GitHub clients

	opsClient := &opslevel.DefaultOpsLevelClient{}
	gitClient := github.NewDefaultGitHubClient(true)

	err := services.ManageServices(opsClient, gitClient)
	if err != nil {
		log.Fatalf("Error managing services: %v", err)
	}

	log.Println("All services processed successfully.")
}
