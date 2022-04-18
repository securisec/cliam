# cliam
## 🚧 WIP

## Description
Cliam is a simple cloud permissions identifier. There are two main components to the CLI.
- `enumerate` which can be used to enumerate specific permissions (recommended)
- `bruteforce` which will try to enumerate all permissions. 

Currently, cliam supports 
- AWS
- GCP

but can be extended with other providers in the future. 


## Usage
Cliam works with credentials obtained from the services well known envars or from passing the commonly required flags from the cli.

```sh
❯❯ cliam --help
Cloud Enumerate is a tool to enumerate cloud credentials for their permissions.

Usage:
  cliam [command]

Available Commands:
  aws         Enumerate AWS credentials for their permissions.
  completion  Generate the autocompletion script for the specified shell
  gcp         Enumerate GCP service accounts for their permissions.
  help        Help about any command

Flags:
  -h, --help   help for cliam

Additional help topics:
  cliam azure      Enumerate Azure credentials for their permissions.

Use "cliam [command] --help" for more information about a command.
```

## Design and extending
In order to persue speed, cliam has the requests hardcoded. 

### AWS
Uses the AWS rest api to make a signed request using the passed in credentials. This greatly adds speed, but makes it a bit more challenging to keep up with adding new permissions. The issue of scale is that AWS uses 3 - 4 variety of requests at the service level

### GCP
Currently, permissions are enumerate using the `cloudresourcemanager` API. This will fail if this service is not enabled, but there are future plans to extend using rest alls to confirm permissions.
