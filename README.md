# gli

_Complete development workflows with GitLab from your terminal_


:warning: *This project is too early in development to be of any practical use to
anyone* :warning:

## tl;dr

## Install

The [just](https://github.com/casey/just) CLI is used to tasks, to install run
the following:

```bash
❯ just install
go install -i -ldflags="-X main.version=loksonarius-dev -s -w"

❯ gli --version
loksonarius-dev
```

### Help

Every command and concept used is documented in the help and usage strings:

```bash
❯ gli help
gli is a CLI to interact with GitLab. The intent is to be more
than a way of running arbitrary CRUD commands against GitLab API
resources, and instead act like an actual developer interface for
typical developer workflows with GitLab.

Usage:
  gli [flags]
  gli [command]

Available Commands:
  cd          Navigate group paths
  help        Help about any command
  login       Log in to a GitLab instance and save it as a target
  ls          List resources and groups for a path
  search      Search for GitLab for resources
  status      Get overview of things
  target      Manage saved GitLab targets

Flags:
  -h, --help   help for gli

Additional help topics:
  gli issue   Interact with GitLab issues

Use "gli [command] --help" for more information about a command.
```
