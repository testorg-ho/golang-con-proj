package opslevel

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

type Service struct {
	Name string `json:"name"`
}

type OpsLevelClient interface {
	GetServices() ([]Service, error)
}

type DefaultOpsLevelClient struct{}

func (c *DefaultOpsLevelClient) GetServices() ([]Service, error) {
	cmd := exec.Command("opslevel", "services", "list", "--json")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	var services []Service
	err = json.Unmarshal(out.Bytes(), &services)
	if err != nil {
		return nil, err
	}

	return services, nil
}
