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
	"fmt"
	"os"
	"path/filepath"

	flag "maunium.net/go/mauflag"
	log "maunium.net/go/maulogger"
)

var configPath = flag.MakeFull("c", "config", "The path to the config file.", "/etc/gh-deployer/config.yaml").String()
var logPath = flag.MakeFull("l", "logs", "The directory to store logs in.", "/var/log/gh-deployer").String()
var debug = flag.MakeFull("d", "debug", "Print debug messages to stdout", "false").Bool()
var wantHelp, _ = flag.MakeHelpFlag()
var config = Config{}

func main() {
	flag.SetHelpTitles(
		"gh-deployer 0.1 - A simple server that listens for changes on GitHub and deploys projects.",
		"gh-deployer [-h] [-c /path/to/config]")

	err := flag.Parse()
	if *wantHelp {
		flag.PrintHelp()
		os.Exit(0)
	} else if err != nil {
		fmt.Println(err)
		flag.PrintHelp()
		os.Exit(1)
	}
	if *debug {
		log.DefaultLogger.PrintLevel = log.LevelDebug.Severity
	}
	log.DefaultLogger.FileMode = 0444
	log.DefaultLogger.FileFormat = func(now string, i int) string {
		return filepath.Join(*logPath, fmt.Sprintf("%s-%02d.log", now, i))
	}

	openConfig()
	startServer()
}
