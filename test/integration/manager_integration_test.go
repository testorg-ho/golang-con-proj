//go:build integration
// +build integration

package integration

import (
	"golang-console-project/opslevel"
	"golang-console-project	git remote add origin <repository-url>/services"
	"testing"
)

type MockOpsLevelClient struct {
	Services []opslevel.Service
}

func (m *MockOpsLevelClient) GetServices() ([]opslevel.Service, error) {
	return m.Services, nil
}

type MockGitHubClient struct {
	CreatedRepos []string
}

func (m *MockGitHubClient) CreateRepo(name string) error {
	m.CreatedRepos = append(m.CreatedRepos, name)
	return nil
}

func TestIntegrationManageServices(t *testing.T) {
	opsClient := &MockOpsLevelClient{
		Services: []opslevel.Service{{Name: "service1"}, {Name: "service2"}},
	}
	gitClient := &MockGitHubClient{}

	err := services.ManageServices(opsClient, gitClient)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(gitClient.CreatedRepos) != 2 {
		t.Errorf("Expected 2 repositories to be created, got %d", len(gitClient.CreatedRepos))
	}
}
