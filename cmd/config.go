package cmd

import (
	"github.com/project-n-oss/interchange/api"
	"github.com/project-n-oss/interchange/app"
)

const cfgEnvPrefix = "INTERCHANGE"

type Config struct {
	Api api.Config `yaml:"Api"`
	App app.Config `yaml:"App"`
}

var DefaultConfig = Config{}
