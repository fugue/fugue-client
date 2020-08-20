# Fugue API Client

This is a command line client for the [Fugue](https://riskmanager.fugue.co/) API.

For CLI documentation and examples, see the [Fugue docs site](https://docs.fugue.co/cli.html).

For information about enabling a client in your account, see
the [API User Guide](https://docs.fugue.co/api.html). More details of the underlying Swagger API are available
[here](https://docs.fugue.co/_static/swagger.html).

This project is under active development and is not yet stable. Commands will
change as we incorporate feedback.

## Install

The easiest way to install the client is to download a prebuilt binary
from [the releases page](https://github.com/fugue/fugue-client/releases).

Place the `fugue` binary in your \$PATH, for example in `/usr/local/bin`.

Change file permissions:

```
chmod 755 /usr/local/bin/fugue
```

For more information, see the [CLI documentation](https://docs.fugue.co/cli.html#installation).

## Environment Variables

The client uses the following _required_ environment variables:

- `FUGUE_API_ID` - your API [client ID](https://docs.fugue.co/api.html#steps)
- `FUGUE_API_SECRET` - your API [client secret](https://docs.fugue.co/api.html#steps)

## Build from Source

Install Go:

```
brew install go
```

Build the client executable:

```
make build
```

Install to \$GOPATH/bin:

```
make install
```

## Usage

Show usage:

```
fugue -h
```

```
Fugue API Client

Usage:
  fugue [command]

Available Commands:
  create      Create a resource
  delete      Delete a resource
  get         Retrieve a resource
  help        Help about any command
  list        List a collection of resources
  scan        Trigger a scan
  sync        Sync files to your account
  update      Update a resource

Flags:
  -h, --help      help for fugue
      --json      outputs the Fugue API JSON response
      --version   version for fugue

Use "fugue [command] --help" for more information about a command.
```

### Command Documentation

- [create](https://docs.fugue.co/cli-create.html)
- [delete](https://docs.fugue.co/cli-delete.html)
- [get](https://docs.fugue.co/cli-get.html)
- [help](https://docs.fugue.co/cli-help.html)
- [list](https://docs.fugue.co/cli-list.html)
- [scan](https://docs.fugue.co/cli-scan.html)
- [sync](https://docs.fugue.co/cli-sync.html)
- [update](https://docs.fugue.co/cli-update.html)

### Aliases

You may use the shorthand `env` instead of `environment` when running commands.

For example:

```
fugue list envs
```

### Debug

To see the HTTP headers and the `json` exchanged between the CLI and the [Fugue](https://riskmanager.fugue.co/) API, set the environment variable `DEBUG=1`. For example:

```
DEBUG=1 fugue list environments
```
