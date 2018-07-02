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

import "strings"

// Reference represents a Git reference.
type Reference string

// IsBranch checks if this reference points to a branch.
func (ref Reference) IsBranch() bool {
	return strings.HasPrefix(string(ref), "refs/heads/")
}

// IsTag checks if this reference points to a tag.
func (ref Reference) IsTag() bool {
	return strings.HasPrefix(string(ref), "refs/tags/")
}

// Name gets the name of the target branch or tag without the refs/type/ prefix
func (ref Reference) Name() string {
	if ref.IsTag() {
		return string(ref[len("refs/tags/"):])
	} else if ref.IsBranch() {
		return string(ref[len("refs/heads/"):])
	}
	return string(ref)
}

// ReferenceType tells what kind of a reference was created/edited/deleted.
type ReferenceType string

// Possible reference types
const (
	ReferenceTypeRepository = "repository"
	ReferenceTypeBranch     = "branch"
	ReferenceTypeTag        = "tag"
)

// Repository is a webhook repository.
type Repository struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Owner         User   `json:"owner"`
	Private       bool   `json:"private"`
	Description   string `json:"description"`
	Fork          bool   `json:"fork"`
	DefaultBranch string `json:"default_branch"`
	MasterBranch  string `json:"master_branch"`
}

// User represents a GitHub user.
type User struct {
	Login      string `json:"login"`
	ID         int    `json:"id"`
	AvatarURL  string `json:"avatar_url"`
	GravatarID string `json:"gravatar_id"`
	Type       string `json:"type"`
	Admin      bool   `json:"site_admin"`
}

// Author is a commit author.
type Author struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Commit contains the metadata of a single commit.
type Commit struct {
	ID        string   `json:"id"`
	TreeID    string   `json:"tree_id"`
	Distinct  bool     `json:"distinct"`
	Message   string   `json:"message"`
	Timestamp string   `json:"timestamp"`
	Author    Author   `json:"author"`
	Committer Author   `json:"committer"`
	Added     []string `json:"added"`
	Removed   []string `json:"removed"`
	Modified  []string `json:"modified"`
}

// PullRequest contains the basic metadata about a pull request.
type PullRequest struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
	State  string `json:"state"`
	Locked bool   `json:"locked"`
	Body   string `json:"body"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ClosedAt  string `json:"closed_at"`
	MergedAt  string `json:"merged_at"`

	MergeCommitHash string `json:"merge_commit_sha"`
}

// Hook contains the metadata for Github webhooks. Used in ping hooks.
type Hook struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Events []string `json:"events"`
	Active bool     `json:"active"`
	Config struct {
		URL         string `json:"url"`
		ContentType string `json:"content_Type"`
	}
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}
