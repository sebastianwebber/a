package task

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Playbook []RunDefinition

func ParsePlaybook(filename string) (*Playbook, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	var out Playbook
	if err := yaml.Unmarshal(f, &out); err != nil {
		return nil, fmt.Errorf("error parsing the yaml file: %w", err)
	}

	return &out, nil
}
