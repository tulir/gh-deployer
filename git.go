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

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func clone(owner, repo, branch string) {
	git.PlainClone(config.GetPath(owner, repo, branch), false, &git.CloneOptions{
		URL:           fmt.Sprintf("https://github.com/%s/%s.git", owner, repo),
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/head/%s", branch)),
	})
}

func remove(owner, repo, branch string) {
	os.RemoveAll(config.GetPath(owner, repo, branch))
}

func pull(owner, repo, branch string) {
	r, _ := git.PlainOpen(config.GetPath(owner, repo, branch))
	r.Pull(&git.PullOptions{})
}
