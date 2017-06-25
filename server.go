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
	"github.com/phayes/hookserve/hookserve"
	log "maunium.net/go/maulogger"
)

func startServer() {
	log.Debugln("Initializing webhook receiver...")
	server := hookserve.NewServer()
	server.Port = config.Port
	server.Secret = config.Secret
	server.Path = config.Path
	server.GoListenAndServe()

	log.Infof("Listening for webhooks on IP:%d%s\n", server.Port, server.Path)
	for event := range server.Events {
		log.Debugf("Received %s event from %s/%s branch %s", event.Type, event.Owner, event.Repo, event.Branch)
		pull(event.Owner, event.Repo, event.Branch)
		run(event.Owner, event.Repo, event.Branch)
	}
}
