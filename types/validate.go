package types

type Validate struct {
	Name   string `yaml:"name"`
	Expect string `yaml:"expect"`
	SQL    string `yaml:"sql"`
}
