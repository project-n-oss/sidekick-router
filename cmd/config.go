package cmd

import (
	"github.com/project-n-oss/sidekick-router/api"
	"github.com/project-n-oss/sidekick-router/app"
)

const cfgEnvPrefix = "SIDEKICKROUTER"

type Config struct {
	Api api.Config `yaml:"Api"`
	App app.Config `yaml:"App"`
}

var DefaultConfig = Config{}
