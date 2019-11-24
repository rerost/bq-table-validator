package types

type Mock struct {
	Table string `yaml:"table"`
	SQL   string `yaml:"sql"`
}

type Validate struct {
	Name   string `yaml:"name"`
	Expect string `yaml:"expect"`
	SQL    string `yaml:"sql"`
	Mocks  []Mock `yaml:"mocks"`
}
