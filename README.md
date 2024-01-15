# Conductor Cli

## Installation -> Go
Conductor commander is a cli app for conductor-oss which is a workflow orchestrator.
The related docs can be found here https://conductor-oss.org
https://github.com/my1795/conductor-cli.git
Install the conductor cli `go install github.com/my1795/conductor-cli@latest`.
Go will automatically install it in your `$GOPATH/bin` directory which should be in your $PATH.

Once installed you should have the `conductor-cli` command available. Confirm by typing `conductor-cli` at a
command line.
## Installation -> Homebrew
```
$ brew tap my1795/tap
$ brew install conductor-cli
```
## Config Set Up
The config file is located at **$HOME/.conductor-cli.yaml** 
the file contains only baseurl at first run of the command
`cat .conductor-cli.yaml
baseurl: http://localhost:8080/api`

if you are using orkes conductor get your key and secret like described https://orkes.io/content/how-to-videos/access-key-and-secret 
and put them in your yaml config like this

```
baseurl: http://localhost:8080/api
key: <<your-key>>
secret <<your-secret>>
```



