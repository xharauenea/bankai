package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	HasSrcDir bool `json:"hasSrcDir"`
	PackageManager string `json:"packageManager"`
	Driver string `json:"driver"`
	Provider string `json:"provider"`
	Packages map[string]bool `json:"packages"`
	Orm string `json:"orm"`
	Auth string `json:"auth"`
	ComponentLib string `json:"componentLib"`
	Alias string `json:"alias"`
}

func CreateConfig(options Config) {
	file, err := json.MarshalIndent(options, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("bankai.config.json", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}