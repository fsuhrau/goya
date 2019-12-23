# goya
goya is a commandline helper connected to jira to easier handle branches / commits

## Requirements
- Go 1.8+
- .goya.yaml in $HOME

## Installation
### form source
``` bash
$ go get -u github.com/fsuhrau/goya
```

### configuration
create a new configuration in $HOME
## create new config
```
touch $HOME/.goya.yaml
open $HOME/.goya.yaml
```

## example config
```
url: https://my.jira.com
username: your_username
password: your_password/api_token
clipboard: true # copy commit message to clipboard
ticket: ([a-zA-Z]+-[0-9]+) # regex to identify ticket number of your current branch
types: # mapping of issue type to branch prefix
  Bug: bugfix/
  Epic: feature/
  User Story: feature/
  Technical Enhancement: feature/
```

## Usage
``` bash
# get informations about ticket and format a branch name to stdout
$ goya branch PROJ-1235

# get informations about ticket and formats a commit message to stdout or clipboard (.goya.yaml)
$ goya commit PROJ-1235

# try to get ticket from current branch, get informations about ticket and formats a commit message to stdout or clipboard (.goya.yaml)
$ goya commit 
```
