# cliam
Multi cloud iam permissions enumeration tool. Currently covers:
- [x] AWS
- [x] GCP
- [x] Azure
- [] Oracle

### 🚧 WIP

[![asciicast](https://asciinema.org/a/goBHd7DlnoOb1x61ljkH2ywc1.png)](https://asciinema.org/a/goBHd7DlnoOb1x61ljkH2ywc1)

## Description
Cliam is a simple cloud permissions identifier. There are two main components to the CLI. Most of the enumerated permissions are list, describe or get permissions. 

- `enumerate` which can be used to enumerate specific permissions (recommended)
- Some service providers have service groups that can check for permissions for a specific subset of services/resources.

## Installation
Download the latest [release](https://github.com/securisec/cliam/releases). DEV tags are current, but not stable.

In order to build the binary locally, cd into the `cli` directory and run `make dev`

## Usage
Cliam works with credentials obtained from the services well known envars or from passing the commonly required flags from the cli.

**It is highly recommond that command completions are set as most of the `enumerate` options have to be specific.** To generate completions, use `cliam completion [shell]` and set according to your shells completion directory.

```
Usage:
  cliam [command]

Available Commands:
  aws         Enumerate AWS credentials for their permissions.
  azure       Enumerate Azure credentials for their permissions.
  completion  Generate the autocompletion script for the specified shell
  firebase    Enumerate Firebase permissions.
  gcp         Enumerate GCP service accounts for their permissions.
  help        Help about any command
  version     Show version and build info

Flags:
  -h, --help                  help for cliam
      --max-threads int       Maximum number of threads to use. (default 5)
      --request-timeout int   Timeout for each request in seconds. (default 5)
      --save-output           Save output to file on success
  -v, --verbose               Enable verbose output.
```

### AWS
Uses the AWS rest api to make a signed request using the passed in credentials. This greatly adds speed, but makes it a bit more challenging to keep up with adding new permissions. The issue of scale is that AWS uses 3 - 4 variety of requests at the service level

Supports obtaining credentials from AWS profile, flags, or default AWS environment variables like `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and optionally `AWS_SESSION_TOKEN`.

```
❯❯ cliam aws                                                        
Enumerate AWS credentials for their permissions.

Usage:
  cliam aws [command]

Available Commands:
  enumerate     Enumerate permissions for specified AWS resources.
  service-group Enumerate permissions for groups of AWS resources.

Flags:
      --access-key-id string         AWS Access Key ID
  -h, --help                         help for aws
      --known-value stringToString   AWS Resource Name. When known-resource-name is set, additional permissions where a resource needs to be specified is enumerated. (default [])
      --profile string               AWS Profile. When profile is set, access-key-id, secret-access-key, and session-token are ignored.
      --region string                AWS Region (default "us-east-1")
      --secret-access-key string     AWS Secret Access Key
      --session-json string          AWS Session JSON file. This flag attempts to read session information from the specified file. Helpful with temporary credentials.
      --session-token string         AWS Session Token

Global Flags:
      --max-threads int       Maximum number of threads to use. (default 5)
      --request-timeout int   Timeout for each request in seconds. (default 5)
  -v, --verbose               Enable verbose output.
```

### Known resources
Cliam for AWS also supports enumerating certain permissions which requires a known value. For instance, when using awscli, we can get a function using `aws lambda get-function --function-name <function_name>`.

This maps directly to cliam where we can use:

```
cliam aws enumerate lamda --known-value function-name=<function_name>
```
This will enumerate all permissions for lambda which takes function-name as a valid argument. This will work with other AWS resources as well.

#### Examples
Bruteforce all serverless resources from an AWS profile
```
❯❯ cliam aws service-group serverless --profile=my-profile
```

Use temporary session tokens obtained to check all ec2 permissions
```
❯❯ cliam aws enumerate ec2 --session-json=creds.json

Where creds.json has
{
  "Type" : "AWS-HMAC",
  "AccessKeyId" : "ASIA...",
  "SecretAccessKey" : "...",
  "Token" : "...",
}
```

Enumerate permissions for s3, iam and ec2
```
❯❯ cliam aws enumerate s3 iam ec2
```

### GCP
GCP supports enumerating from a specific service account json file. ~Credentials from the GCP environment variables `GOOGLE_APPLICATION_CREDENTIALS` and `CLOUDSDK_CORE_PROJECT` are also supported~.

Because there are two modes for GCP enumeration, use enumerate if `cloudresourcemanager` is enabled or use `rest` to enumerate specific permissions.


```
cliam gcp --help
Enumerate GCP service accounts for their permissions.

Usage:
  cliam gcp [command]

Available Commands:
  bruteforce  Enumerate all GCP permissions
  enumerate   Enumerate specified GCP permissions
  rest        GCP permissions using the REST API

Flags:
  -h, --help                     help for gcp
      --project-id string        GCP project id
      --region string            GCP Region (default "us-central1")
      --service-account string   GCP service account path
      --zone string              GCP Zone (default "us-central1-a")

Global Flags:
      --max-threads int       Maximum number of threads to use. (default 5)
      --request-timeout int   Timeout for each request in seconds. (default 10)

Use "cliam gcp [command] --help" for more information about a command.
```

### Azure
Azure enumeration supports various authentication methods inlcuding service principals (via client id and secret), certificate MSI (via MSI token), or default auth. It also supports all the common azure environment variables 
- AZURE_CLIENT_ID : client id or usrname
- AZURE_CLIENT_SECRET : client secret
- AZURE_SUBSCRIPTION_ID : subscription id
- AZURE_CLIENT_CERTIFICATE_PATH : path to certificate
- AZURE_TENANT_ID : tenant id
- CLIAM_AZURE_OAUTH_TOKEN : existing oauth token

```
❯❯ cliam azure
Enumerate Azure credentials for their permissions.
In most cases, a valid Azure Subscription ID is required. If one 
is not provided, the CLI will attempt to lookup available subscriptions 
and use the first one.

Usage:
  cliam azure [command]

Available Commands:
  access-token  Get access token to use with REST apis
  enumerate     Enumerate permissions for specified azure resources.
  service       Enumerate permissions for a specific group of azure services.
  subscriptions List all subscriptions
```

## Debug
cliam supports two environment variables to show debug output
- DEBUG=true (shows status codes of requests)
- VERBOSE=true (shows body of requests)
