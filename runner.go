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
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// RunnerConfig contains the branch-specific deployment instructions.
type RunnerConfig struct {
	Directory   string   `yaml:"-"`
	Shell       string   `yaml:"shell"`
	ShellArgs   []string `yaml:"shell-args"`
	Environment []string `yaml:"env"`
	Commands    []string `yaml:"commands"`
}

func run(owner, repo, branch string) {
	dir := config.GetPath(owner, repo, branch)
	dat, err := ioutil.ReadFile(filepath.Join(dir, ".gh-deployer.yaml"))
	if err != nil {
		fmt.Printf("Failed to read deployer run config of %s/%s branch %s: %s\n", owner, repo, branch, err)
		return
	}

	runConfig := RunnerConfig{}
	err = yaml.Unmarshal(dat, &runConfig)
	if err != nil {
		fmt.Printf("Failed to parse deployer run config of %s/%s branch %s: %s\n", owner, repo, branch, err)
		return
	}
	runConfig.Directory = dir
	go runConfig.run()
}

func (rconf RunnerConfig) run() {
	stdoutFile, _ := os.Create(filepath.Join(rconf.Directory, ".stdout"))
	stderrFile, _ := os.Create(filepath.Join(rconf.Directory, ".stderr"))
	stdoutWriter := bufio.NewWriter(stdoutFile)
	stderrWriter := bufio.NewWriter(stderrFile)
	infoMessages := io.MultiWriter(stdoutWriter, stderrWriter)

	for _, rawCommand := range rconf.Commands {
		var command string
		var args []string
		if len(rconf.Shell) > 0 {
			command = rconf.Shell
			args = append(rconf.ShellArgs, rawCommand)
		} else {
			parts := strings.Split(rawCommand, " ")
			command = parts[0]
			if len(parts) > 1 {
				args = parts[1:]
			}
		}
		fmt.Fprintln(infoMessages, "--------------------------------------------------")
		fmt.Fprintln(infoMessages, "[gh-deployer] Preparing command", rawCommand)

		cmd := exec.Command(command, args...)
		cmd.Dir = rconf.Directory
		cmd.Env = append(os.Environ(), rconf.Environment...)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintf(stderrWriter, "[gh-deployer] Failed to open stdout of command: %s!\n", err)
			stdout = nil
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Fprintf(stderrWriter, "[gh-deployer] Failed to open stderr of command: %s!\n", err)
			stderr = nil
		}

		err = cmd.Start()
		if err != nil {
			fmt.Fprintf(stderrWriter, "[gh-deployer] Failed to execute command: %s!\n", err)
			continue
		}
		fmt.Fprintln(infoMessages, "[gh-deployer] Command started. Piping output...")

		if stdout != nil {
			go io.Copy(stdoutWriter, stdout)
		}
		if stderr != nil {
			go io.Copy(stderrWriter, stderr)
		}
		err = cmd.Wait()
		fmt.Fprintf(stderrWriter, "[gh-deployer] Error while waiting for command: %s!\n", err)
		fmt.Fprintln(infoMessages, "[gh-deployer] Command execution finished.")
	}
}
