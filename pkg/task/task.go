package task

type RunDefinition struct {
	Hosts       string   `yaml:"hosts"`
	Become      bool     `yaml:"become"`
	GatherFacts bool     `yaml:"gather_facts,omitempty"`
	Roles       []Roles  `yaml:"roles,omitempty"`
	Name        string   `yaml:"name,omitempty"`
	BecomeUser  string   `yaml:"become_user,omitempty"`
	VarsFiles   []string `yaml:"vars_files,omitempty"`
	Tasks       []Tasks  `yaml:"tasks,omitempty"`
}
type Roles struct {
	Name string `yaml:"name"`
}
type Yum struct {
	Name  string `yaml:"name"`
	State string `yaml:"state"`
}
type Service struct {
	Name  string `yaml:"name"`
	State string `yaml:"state"`
}
type Tasks struct {
	Name    string  `yaml:"name"`
	Yum     Yum     `yaml:"yum,omitempty"`
	Service Service `yaml:"service,omitempty"`
}
