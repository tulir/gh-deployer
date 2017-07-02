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
	"maunium.net/go/githuuk"
	log "maunium.net/go/maulogger"
)

func startServer() {
	log.Debugln("Initializing webhook receiver...")
	server := githuuk.NewServer()
	server.Host = config.Host
	server.Port = config.Port
	server.Secret = config.Secret
	server.Path = config.Path
	server.AsyncListenAndServe()

	log.Infof("Listening for webhooks on %s:%d%s\n", server.Host, server.Port, server.Path)
	for rawEvent := range server.Events {
		switch evt := rawEvent.(type) {
		case *githuuk.PushEvent:
			log.Debugf("%s pushed to %s branch %s\n", evt.Sender.Login, evt.Repository.FullName, evt.Ref.Name())
			pull(evt.Repository.Owner.Login, evt.Repository.Name, evt.Ref.Name())
			run(evt.Repository.Owner.Login, evt.Repository.Name, evt.Ref.Name())
		case *githuuk.DeleteEvent:
			log.Debugf("%s deleted branch %s of %s", evt.Sender.Login, evt.Ref.Name(), evt.Repository.FullName)
			remove(evt.Repository.Owner.Login, evt.Repository.Name, evt.Ref.Name())
		}
	}
}
