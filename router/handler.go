package router

import (
	"encoding/json"
	"net/http"

	"github.com/hfantin/goblank/autoconfig"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	build := make(map[string]interface{})
	status := make(map[string]interface{})
	status["group"] = autoconfig.Env.Group
	status["name"] = autoconfig.Env.Name
	status["version"] = autoconfig.Env.Version
	status["time"] = currentTime
	build["build"] = status
	sendJson(w, http.StatusOK, build)

}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	sendJson(w, http.StatusOK, map[string]string{"health": "OK"})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	sendJson(w, http.StatusNotFound, map[string]string{"msg": "Not Found"})
}

func sendError(w http.ResponseWriter, message string) {
	sendJson(w, http.StatusInternalServerError, map[string]string{"error": message})
}

func sendJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if payload != nil {
		response, _ := json.Marshal(payload)
		w.Write(response)
	}
}
