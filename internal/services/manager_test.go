package services

import (
	"errors"
	"testing"

	"golang-console-project/internal/github"
	"golang-console-project/internal/opslevel"
	"golang-console-project/internal/opslevel/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestManageServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockOpsClient := mocks.NewMockOpsLevelClient(ctrl)
	mockGitClient := github.NewMockGitHubClient(ctrl)

	services := []opslevel.Service{
		{Name: "service1"},
		{Name: "service2"},
	}

	// Mock opsClient.GetServices to return the services list
	mockOpsClient.EXPECT().GetServices().Return(services, nil).Times(1)

	// Mock gitClient.CreateRepo to be called for each service
	mockGitClient.EXPECT().CreateRepo("service1").Return(nil).Times(1)
	mockGitClient.EXPECT().CreateRepo("service2").Return(errors.New("repo creation failed")).Times(1)

	err := ManageServices(mockOpsClient, mockGitClient)

	// Assert no error is returned from ManageServices
	assert.NoError(t, err)
}
