package github

import (
	"fmt"
	"os"
	"os/exec"
)

type GitHubClient interface {
	CreateRepo(name string) error
}

type DefaultGitHubClient struct {
	createRepos bool
	Username    string
	Password    string
	URL         string
}

func NewDefaultGitHubClient(createRepos bool) *DefaultGitHubClient {
	username := os.Getenv("GITHUB_USERNAME")
	password := os.Getenv("GITHUB_PASSWORD")
	url := os.Getenv("GITHUB_URL")
	return &DefaultGitHubClient{
		createRepos: createRepos,
		Username:    username,
		Password:    password,
		URL:         url,
	}
}

func (c *DefaultGitHubClient) CreateRepo(name string) error {
	if !c.createRepos {
		fmt.Printf("Simulating creation of repository: %s\n", name)
		return nil
	}

	// Check if the repository exists
	cmd := exec.Command("gh", "repo", "view", name)
	if err := cmd.Run(); err != nil {
		// If the repository does not exist, create it
		cmd = exec.Command("gh", "repo", "create", name, "--public", "--confirm")
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
