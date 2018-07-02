// githuuk - A GitHub webhook receiver written in Go.
// Copyright (C) 2017 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package githuuk

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Version is the version of Hookserve.
var Version = "0.1.0"

// Server is the main server container.
type Server struct {
	Host     string
	Port     uint16
	Path     string
	PingPath string
	Secret   string
	Events   chan Event
}

// NewServer creates a webhook server with sensible defaults.
// Default settings:
//   Port: 80
//   Path: /webhook
//   Ping path: /webhook/ping
//   Ignore tags: true
func NewServer() *Server {
	return &Server{
		Port:     80,
		Path:     "/webhook",
		PingPath: "/webhook/ping",
		Events:   make(chan Event, 10),
	}
}

// ListenAndServe runs the server and returns if an error occurs.
func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", s.Host, s.Port), s)
}

// AsyncListenAndServe runs the server inside a Goroutine and panics if an error occurs.
func (s *Server) AsyncListenAndServe() {
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
}

// CheckSignature checks that the request is from GitHub.
func (s *Server) CheckSignature(w http.ResponseWriter, r *http.Request, body []byte) bool {
	if len(s.Secret) > 0 {
		sig := r.Header.Get("X-Hub-Signature")

		if len(sig) == 0 {
			http.Error(w, "Forbidden - Missing X-Hub-Signature", http.StatusForbidden)
			return false
		}

		mac := hmac.New(sha1.New, []byte(s.Secret))
		mac.Write(body)
		expectedMAC := mac.Sum(nil)
		expectedSig := fmt.Sprintf("sha1=%s", hex.EncodeToString(expectedMAC))
		if !hmac.Equal([]byte(expectedSig), []byte(sig)) {
			http.Error(w, "Forbidden - X-Hub-Signature verification failed", http.StatusForbidden)
			return false
		}
	}
	return true
}

// ServeHTTP implements the http.Handler interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method == "GET" && r.URL.Path == s.PingPath {
		w.Write([]byte("OK"))
		return
	} else if r.URL.Path != s.Path {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	} else if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	eventType := EventType(r.Header.Get("X-GitHub-Event"))
	if eventType == "" {
		http.Error(w, "Bad Request - Missing X-GitHub-Event Header", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !s.CheckSignature(w, r, body) {
		return
	}

	event, status, err := s.ParseEvent(eventType, body)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	s.Events <- event

	w.Header().Set("Server", fmt.Sprintf("githuuk/%s", Version))
	w.Write([]byte("{}"))
}

// ParseEvent parses a JSON body into an event struct of the given type.
func (s *Server) ParseEvent(eventType EventType, body []byte) (Event, int, error) {
	var event Event
	base := BaseEvent{Type: eventType}
	switch eventType {
	case EventPush:
		event = &PushEvent{BaseEvent: base}
	case EventPullRequest:
		event = &PullRequestEvent{BaseEvent: base}
	case EventPing:
		event = &PingEvent{BaseEvent: base}
	case EventCreate:
		event = &CreateEvent{BaseEvent: base}
	case EventDelete:
		event = &DeleteEvent{BaseEvent: base}
	default:
		return nil, http.StatusNotImplemented, fmt.Errorf("501 Not Implemented - Unknown event type %s", eventType)
	}
	err := json.Unmarshal(body, event)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return event, http.StatusOK, nil
}
