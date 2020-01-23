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

Every command and concept used is document in help and usage strings:

```bash
❯ gli -h
gli is a CLI to interact with GitLab. The intent is to be more
than a way of running arbitrary CRUD commands against GitLab API
resources, and act like an actual developer interface for typical
developer workflows with GitLab.

Usage:
  gli [flags]
  gli [command]

Available Commands:
  group       Navaigate and activate groups
  help        Help about any command
  search      Search for GitLab for resources

Flags:
  -h, --help            help for gli
  -t, --target string   target GitLab instance to run commands against
      --version         version for gli
```
