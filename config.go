package main

import (
	"github.com/BurntSushi/toml"
	"github.com/ismdeep/args"
	"io/ioutil"
	"os"
)

type config struct {
	Name   string
	Images []struct {
		Name string
		Tags []string
	}
}

var Config *config

func init() {
	path := "./config.toml"
	if os.Getenv("DOCKER_IMAGES_DUMPER_CONFIG") != "" {
		path = os.Getenv("DOCKER_IMAGES_DUMPER_CONFIG")
	}
	if args.Exists("-c") {
		path = args.GetValue("-c")
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	Config = &config{}
	if err := toml.Unmarshal(content, Config); err != nil {
		panic(err)
	}
}
