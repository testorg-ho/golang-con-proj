package main

import (
	"errors"
	"golang-console-project/internal/opslevel"
	"golang-console-project/internal/services"
	"testing"
)

type MockOpsLevelClient struct {
	Services []opslevel.Service
	Err      error
}

func (m *MockOpsLevelClient) GetServices() ([]opslevel.Service, error) {
	return m.Services, m.Err
}

type MockGitHubClient struct {
	CreatedRepos []string
	Err          error
}

func (m *MockGitHubClient) CreateRepo(name string) error {
	if m.Err != nil {
		return m.Err
	}
	m.CreatedRepos = append(m.CreatedRepos, name)
	return nil
}

func TestMainSuccess(t *testing.T) {
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

func TestMainOpsLevelError(t *testing.T) {
	opsClient := &MockOpsLevelClient{
		Err: errors.New("OpsLevel error"),
	}
	gitClient := &MockGitHubClient{}

	err := services.ManageServices(opsClient, gitClient)
	if err == nil {
		t.Fatal("Expected an error but got nil")
	}
}

func TestMainGitHubError(t *testing.T) {
	opsClient := &MockOpsLevelClient{
		Services: []opslevel.Service{{Name: "service1"}},
	}
	gitClient := &MockGitHubClient{
		Err: errors.New("GitHub error"),
	}

	err := services.ManageServices(opsClient, gitClient)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(gitClient.CreatedRepos) != 0 {
		t.Errorf("Expected 0 repositories to be created, got %d", len(gitClient.CreatedRepos))
	}
}
