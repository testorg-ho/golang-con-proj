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
}

func NewDefaultGitHubClient() *DefaultGitHubClient {
	createRepos := os.Getenv("CREATE_REPOS") == "true"
	return &DefaultGitHubClient{createRepos: createRepos}
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
