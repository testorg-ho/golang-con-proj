//go:build unit
// +build unit

package services

import (
	"errors"
	"golang-console-project/internal/opslevel"
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

func TestManageServices(t *testing.T) {
	mockOpsClient := &MockOpsLevelClient{
		Services: []opslevel.Service{{Name: "service1"}, {Name: "service2"}},
	}
	mockGitClient := &MockGitHubClient{}

	err := ManageServices(mockOpsClient, mockGitClient)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(mockGitClient.CreatedRepos) != 2 {
		t.Errorf("Expected 2 repositories to be created, got %d", len(mockGitClient.CreatedRepos))
	}
}

func TestManageServicesWithError(t *testing.T) {
	mockOpsClient := &MockOpsLevelClient{
		Services: []opslevel.Service{{Name: "service1"}},
	}
	mockGitClient := &MockGitHubClient{
		Err: errors.New("GitHub error"),
	}

	err := ManageServices(mockOpsClient, mockGitClient)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(mockGitClient.CreatedRepos) != 0 {
		t.Errorf("Expected 0 repositories to be created, got %d", len(mockGitClient.CreatedRepos))
	}
}
