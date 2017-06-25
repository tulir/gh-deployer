# GitHub deployer
A simple server that listens for changes on GitHub and deploys projects.

## Setup
0. Have a functional [Go installation](https://golang.org/doc/install)
1. Download and compile gh-deployer using `go get maunium.net/go/gh-deployer`
2. Configure and start up gh-deployer ([example config](https://github.com/tulir/gh-deployer/blob/master/example-config.yaml)). The default config path is `/etc/gh-deployer/config.yaml`.
3. Configure Github webhooks according to your gh-deployer config.
4. Create `.gh-deployer.yaml` in the root of the repository to deploy ([example deploy config](https://github.com/tulir/gh-deployer/blob/master/example-runner.yaml)). If you have gh-deployer started and Github webhooks set up, the server should run the commands as soon as you push the deploy config.

Compiled builds coming soon™.
