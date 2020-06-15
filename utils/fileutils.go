package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadJson(filename string) map[string]string {
	jsonFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to read %s file: %v\n", filename, err)
		return nil
	}
	var jsonMap map[string]string
	err = json.Unmarshal([]byte(jsonFile), &jsonMap)
	if err != nil {
		log.Printf("Failed to parse %s file: %v\n", filename, err)
		return nil
	}
	return jsonMap
}
