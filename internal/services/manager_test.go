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
	mockGitClient.On("CreateRepo", "service1").Return(nil).Once()
	mockGitClient.On("CreateRepo", "service2").Return(errors.New("repo creation failed")).Once()

	// Call the function under test
	err := ManageServices(mockOpsClient, mockGitClient)

	// Assert expectations
	assert.NoError(t, err)
	mockOpsClient.AssertExpectations(t)
	mockGitClient.AssertExpectations(t)

	// Verify the number of calls
	mockOpsClient.AssertNumberOfCalls(t, "GetServices", 1)
	mockGitClient.AssertNumberOfCalls(t, "CreateRepo", 2)

	// Verify CreateRepo called with expected arguments in sequence
	mockGitClient.AssertCalled(t, "CreateRepo", "service1")
	mockGitClient.AssertCalled(t, "CreateRepo", "service2")
	mockGitClient.AssertCalled(t, "CreateRepo", mock.MatchedBy(func(arg string) bool {
		return arg == "service1"
	}))
	mockGitClient.AssertCalled(t, "CreateRepo", mock.MatchedBy(func(arg string) bool {
		return arg == "service2"
	}))
}


func TestManageServicesParam(t *testing.T) {
    tests := []struct {
        name          string
        services      []opslevel.Service
        mockGitClient func() *MockGitHubClient
        expectedError error
    }{
        {
            name: "All repositories created successfully",
            services: []opslevel.Service{
                {Name: "service1"},
                {Name: "service2"},
            },
            mockGitClient: func() *MockGitHubClient {
                mockGitClient := new(MockGitHubClient)
                mockGitClient.On("CreateRepo", "service1").Return(nil).Once()
                mockGitClient.On("CreateRepo", "service2").Return(nil).Once()
                return mockGitClient
            },
            expectedError: nil,
        },
        {
            name: "One repository creation fails",
            services: []opslevel.Service{
                {Name: "service1"},
                {Name: "service2"},
            },
            mockGitClient: func() *MockGitHubClient {
                mockGitClient := new(MockGitHubClient)
                mockGitClient.On("CreateRepo", "service1").Return(nil).Once()
                mockGitClient.On("CreateRepo", "service2").Return(errors.New("repo creation failed")).Once()
                return mockGitClient
            },
            expectedError: nil, // ManageServices might handle errors internally
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockOpsClient := new(MockOpsLevelClient)
            mockOpsClient.On("GetServices").Return(tt.services, nil)

            mockGitClient := tt.mockGitClient()

            // Call the function under test
            err := ManageServices(mockOpsClient, mockGitClient)

            // Assert expectations
            assert.Equal(t, tt.expectedError, err)
            mockOpsClient.AssertExpectations(t)
            mockGitClient.AssertExpectations(t)
        })
    }
}
