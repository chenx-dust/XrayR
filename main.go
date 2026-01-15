package main

import (
	"net/http"
	_ "net/http/pprof"

	log "github.com/sirupsen/logrus"

	"github.com/XrayR-project/XrayR/cmd"
)

func main() {
	go func() {
		http.ListenAndServe("[::]:6060", nil)
	}()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
