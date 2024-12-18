package config

type Option struct {
	Label string
	Value string
	Usage int
}

type Patch struct {
	Prefix string
	Major  int
	Minor  int
	Patch  int
	Suffix string
}
