package gitmux

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/Chaitanyabsprip/gitmux/tmux"
)

// Config configures output formatting.
type Config struct{ Tmux tmux.Config }

// default config (decoded in init)
var DefaultCfg Config

//go:embed .gitmux.yml
var CfgBytes []byte

func init() {
	if err := yaml.Unmarshal(CfgBytes, &DefaultCfg); err != nil {
		panic(fmt.Sprintf("default config is invalid: %v", err))
	}
}
