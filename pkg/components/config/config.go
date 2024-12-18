package config

type Option struct {
	Label string
	Value string
	Usage int
}

type Config struct {
	Options []Option
	Patch   int8
}
