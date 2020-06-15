package config

import (
	"time"

	"github.com/hfantin/goblank/utils"
	"github.com/namsral/flag"
)

type environment struct {
	ServerPort         int
	InicializationTime time.Time
	BuildFile          string
	Group              string
	Version            string
	Name               string
}

var Env *environment

func Load() {
	Env = &environment{}
	flag.IntVar(&Env.ServerPort, "SERVER_PORT", 5000, "This is the server port.")
	flag.StringVar(&Env.BuildFile, "BUILD_FILE", "build.json", "This is the build file.")
	flag.Parse()
	Env.InicializationTime = time.Now()
	buildJson := utils.ReadJson(Env.BuildFile)
	if buildJson != nil {
		Env.Group = buildJson["group"]
		Env.Name = buildJson["name"]
		Env.Version = buildJson["version"]
	}
}
