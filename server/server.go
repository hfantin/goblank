package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hfantin/goblank/config"
	"github.com/hfantin/goblank/router"
)

func New() *http.Server {
	servidor := &http.Server{
		Handler:      router.New(),
		Addr:         fmt.Sprintf(":%d", config.Env.ServerPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return servidor
}
