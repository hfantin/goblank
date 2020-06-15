package autoconfig

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/namsral/flag"
)

type environment struct {
	ServerPort int
	Group      string
	Version    string
	Name       string
}

var Env *environment

func init() {
	Env = &environment{}
	flag.IntVar(&Env.ServerPort, "SERVER_PORT", 5000, "This is the server port.")
	flag.Parse()
	buildJson := readBuildFile()
	if buildJson != nil {
		Env.Group = buildJson["group"]
		Env.Name = buildJson["name"]
		Env.Version = buildJson["version"]
	}
}

func readBuildFile() map[string]string {
	jsonFile, err := ioutil.ReadFile("build.json")
	if err != nil {
		log.Println("Failed to read build.json file:", err)
		return nil
	}
	var jsonMap map[string]string
	err = json.Unmarshal([]byte(jsonFile), &jsonMap)
	if err != nil {
		log.Println("Failed to parse build.json file", err)
		return nil
	}
	return jsonMap
}
