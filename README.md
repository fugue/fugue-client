# Fugue API Client

This is a command line client for the Fugue API.

For information about enabling a client in your account, see
[the documenation](https://riskmanagerdocs.fugue.co/api.html). More details of the underlying Swagger API are available
[here](https://riskmanagerdocs.fugue.co/Resources/API/swagger.html).

This project is under active development and is not yet stable. Commands will
change as we incorporate feedback.

## Environment Variables

The client uses the following *required* environment variables:

* `FUGUE_API_ID` - your API client ID
* `FUGUE_API_SECRET` - your API client secret

## Build

Build the client executable:

```
make build
```

Install to $GOPATH/bin:

```
make install
```

## Usage

Show usage: 

```
fugue -h
```

Show usage for a subcommand:

```
fugue list -h
```

List environments:

```
fugue list environments
```

List environment scans:

```
fugue list scans [environment-id]
```

List environment events:

```
fugue list events [environment-id]
```

Trigger a scan:

```
fugue scan [environment-id]
```

Compliance by resource types:

```
fugue get compliance-by-resource-types [scan-id]
```

Compliance by rules:

```
fugue get compliance-by-rules [scan-id]
```

## Aliases

You may use the shorthand `env` instead of `environment` when running commands.

For example:

```
fugue list envs
```
