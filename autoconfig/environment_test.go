package autoconfig_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hfantin/goblank/autoconfig"
)

func init() {
	os.Setenv("SERVER_PORT", "9999")
	os.Setenv("BUILD_FILE", "../build.json")
	fmt.Println("mock")
}
func TestInit(t *testing.T) {

	t.Run("check the default environment configuration", func(t *testing.T) {
		env := autoconfig.Env
		if env.ServerPort != 5000 {
			t.Errorf("Default port is different from 5000")
		}
		if env.BuildFile != "build.json" {
			t.Errorf("Default build file is different from build.json")
		}
	})

}
