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

// EventType represents a GitHub webhook event.
type EventType string

// Supported event types
const (
	EventPush        EventType = "push"
	EventPullRequest EventType = "pull_request"
	EventCreate      EventType = "create"
	EventDelete      EventType = "delete"
	EventPing        EventType = "ping"
)

// Event is an abstract event.
type Event interface {
	GetType() EventType
}

// BaseEvent contains the base fields for all Events.
type BaseEvent struct {
	Type       EventType  `json:"-"`
	Repository Repository `json:"repository"`
	Sender     User       `json:"sender"`
}

// GetType returns the type of this event.
func (evt *BaseEvent) GetType() EventType {
	return evt.Type
}

// PingEvent contains the data for the webhook creation ping.
type PingEvent struct {
	BaseEvent
	Zen    string `json:"zen"`
	HookID int    `json:"hook_id"`
	Hook   Hook   `json:"hook"`
}

// PushEvent contains the data in a push-type webhook event.
type PushEvent struct {
	BaseEvent
	Commits    []Commit  `json:"commits"`
	HeadCommit Commit    `json:"head_commit"`
	Ref        Reference `json:"ref"`
	Pusher     User      `json:"pusher"`
	Deleted    bool      `json:"deleted"`
	Created    bool      `json:"created"`
	Forced     bool      `json:"forced"`
}

// PullRequestEvent contains the data in a pull_request-type webhook event.
type PullRequestEvent struct {
	BaseEvent
	Action          string `json:"action"`
	NumberOfChanges int    `json:"number"`
}

// CreateEvent contains the data in a create (branch/tag) webhook event.
type CreateEvent struct {
	BaseEvent
	RefType      ReferenceType `json:"ref_type"`
	Ref          Reference     `json:"ref"`
	MasterBranch string        `json:"master_branch"`
	Description  string        `json:"description"`
}

// DeleteEvent contains the data in a delete (branch/tag) webhook event.
type DeleteEvent struct {
	BaseEvent
	RefType ReferenceType `json:"ref_type"`
	Ref     Reference     `json:"ref"`
}
