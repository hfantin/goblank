package autoconfig

import (
	"github.com/hfantin/goblank/utils"
	"github.com/namsral/flag"
)

type environment struct {
	ServerPort int
	BuildFile  string
	Group      string
	Version    string
	Name       string
}

var Env *environment

func init() {
	Env = &environment{}
	flag.IntVar(&Env.ServerPort, "SERVER_PORT", 5000, "This is the server port.")
	flag.StringVar(&Env.BuildFile, "BUILD_FILE", "build.json", "This is the build file.")
	flag.Parse()
	buildJson := utils.ReadJson(Env.BuildFile)
	if buildJson != nil {
		Env.Group = buildJson["group"]
		Env.Name = buildJson["name"]
		Env.Version = buildJson["version"]
	}
}
