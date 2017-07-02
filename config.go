// gh-deployer - A simple server that listens for changes on GitHub and deploys projects.
// Copyright (C) 2017 Tulir Asokan

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"

	log "maunium.net/go/maulogger"
)

// Config is the main global config struct.
type Config struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	Path          string `yaml:"path"`
	Secret        string `yaml:"secret"`
	PullDirectory string `yaml:"pull-directory"`
}

// GetPath gets the path to a pull directory
func (config Config) GetPath(owner, repo, branch string) (str string) {
	str = strings.Replace(config.PullDirectory, "$REPO_NAME", repo, -1)
	str = strings.Replace(str, "$REPO_OWNER", owner, -1)
	str = strings.Replace(str, "$BRANCH", branch, -1)
	return
}

func openConfig() {
	log.Debugln("Opening config from", *configPath)
	data, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatalln("Failed to read config:", err)
		os.Exit(2)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalln("Failed to parse config:", err)
		os.Exit(3)
	}
}
