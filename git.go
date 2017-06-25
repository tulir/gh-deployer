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
	fmt.Printf("Cloning %s/%s branch %s\n", owner, repo, branch)
	r, err := git.PlainClone(config.GetPath(owner, repo, branch), false, &git.CloneOptions{
		URL:           fmt.Sprintf("https://github.com/%s/%s.git", owner, repo),
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
	})
	if err != nil {
		fmt.Println("Failed to clone:", err)
		return
	}
	r.Pull(&git.PullOptions{})
}

func remove(owner, repo, branch string) {
	path := config.GetPath(owner, repo, branch)
	fmt.Println("Removing", path)
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("Failed to remove:", err)
	}
}

func pull(owner, repo, branch string) {
	fmt.Printf("Pulling %s/%s branch %s\n", owner, repo, branch)
	path := config.GetPath(owner, repo, branch)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
	r, err := git.PlainOpen(path)
	if err != nil {
		fmt.Println("Failed to open repo at", path)
		remove(owner, repo, branch)
		clone(owner, repo, branch)
		return
	}
	r.Pull(&git.PullOptions{})
}
