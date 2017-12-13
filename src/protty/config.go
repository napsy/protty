package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gotk3/gotk3/glib"
)

type config struct {
	sessions map[string]map[string]string
}

var cfg config

func loadConfiguration() error {

	cfgDir := glib.GetUserConfigDir()

	b, err := ioutil.ReadFile(fmt.Sprintf("%s/protty-config.json", cfgDir))
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &cfg)
}

func saveConfiguration() error {
	for e := sessions.Front(); e != nil; e = e.Next() {
		s := e.Value.(session)
		settings := s.toMap()
		cfg.sessions[s.name()] = settings
	}
	b, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}
	cfgDir := glib.GetUserConfigDir()

	return ioutil.WriteFile(fmt.Sprintf("%s/protty-config.json", cfgDir), b, 0666)
}
