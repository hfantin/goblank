package config_test

import (
	"os"
	"testing"

	"github.com/hfantin/goblank/config"
)

func TestLoad(t *testing.T) {
	t.Run("check the default environment configuration", func(t *testing.T) {
		os.Setenv("SERVER_PORT", "9999")
		os.Setenv("BUILD_FILE", "../build.json")
		config.Load()
		if config.Env.ServerPort != 9999 {
			t.Errorf("Default port is different from 9999")
		}
		if config.Env.BuildFile != "../build.json" {
			t.Errorf("Default build file is different from build.json")
		}
	})
}
