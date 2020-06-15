package utils_test

import (
	"testing"

	"github.com/hfantin/goblank/utils"
)

func TestReadJson(t *testing.T) {
	t.Run("returns file info as a map", func(t *testing.T) {
		json := utils.ReadJson("../build.json")
		if json == nil {
			t.Errorf("json is null")
		}
	})
}
