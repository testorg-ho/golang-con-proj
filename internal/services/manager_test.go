package services

import (
	"errors"
	"testing"

	"golang-console-project/internal/opslevel"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockOpsLevelClient is a mock implementation of the OpsLevelClient interface.
type MockOpsLevelClient struct {
	mock.Mock
}

func (m *MockOpsLevelClient) GetServices() ([]opslevel.Service, error) {
	args := m.Called() 
	return args.Get(0).([]opslevel.Service), args.Error(1)
}

// MockGitHubClient is a mock implementation of the GitHubClient interface.
type MockGitHubClient struct {
	mock.Mock
}

func (m *MockGitHubClient) CreateRepo(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func TestManageServices(t *testing.T) {
	mockOpsClient := new(MockOpsLevelClient)
	mockGitClient := new(MockGitHubClient)

	services := []opslevel.Service{
		{Name: "service1"},
		{Name: "service2"},
	}

	// Set up expectations for the mock clients
	mockOpsClient.On("GetServices").Return(services, nil)
	mockGitClient.On("CreateRepo", "service1").Return(nil)
	mockGitClient.On("CreateRepo", "service2").Return(errors.New("repo creation failed"))

	// Call the function under test
	err := ManageServices(mockOpsClient, mockGitClient)

	// Assert expectations
	assert.NoError(t, err)
	mockOpsClient.AssertExpectations(t)
	mockGitClient.AssertExpectations(t)
}
