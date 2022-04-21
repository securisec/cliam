# cliam
Multi cloud iam permissions enumeration tool. Currently covers:
[x] AWS
[x] GCP
[] Azure
[] Oracle

### 🚧 WIP (bugs are expected)

[![asciicast](https://asciinema.org/a/goBHd7DlnoOb1x61ljkH2ywc1.png)](https://asciinema.org/a/goBHd7DlnoOb1x61ljkH2ywc1)

## Description
Cliam is a simple cloud permissions identifier. There are two main components to the CLI. Most of the enumerated permissions are list, describe or get permissions. Only permissions that does not require a specific resource are tested.

- `enumerate` which can be used to enumerate specific permissions (recommended)
- Some service providers have service groups that can check for permissions for a specific subset of services/resources.


## Usage
Cliam works with credentials obtained from the services well known envars or from passing the commonly required flags from the cli.

**It is highly recommond that command completions are set as most of the `enumerate` options have to be specific.** To generate completions, use `cliam completion [shell]` and set according to your shells completion directory.

```
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

Supports obtaining credentials from AWS profile, flags, or default AWS environment variables like `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and optionally `AWS_SESSION_TOKEN`.

```
cliam aws --help
Enumerate AWS credentials for their permissions.

Usage:
  cliam aws [command]

Available Commands:
  common      Enumerate permissions for common AWS resources.
  compute     Enumerate permissions for common compute AWS resources.
  databases   Enumerate permissions for common AWS database resources.
  enumerate   Enumerate permissions for specified AWS resources.
  serverless  Enumerate permissions for common serverless AWS resources.
  storage     Enumerate permissions for common storage AWS resources.

Flags:
      --access-key-id string       AWS Access Key ID
  -h, --help                       help for aws
      --profile string             AWS Profile. When profile is set, access-key-id, secret-access-key, and session-token are ignored.
      --region string              AWS Region (default "us-east-1")
      --secret-access-key string   AWS Secret Access Key
      --session-token string       AWS Session Token

Global Flags:
      --max-threads int       Maximum number of threads to use. (default 5)
      --request-timeout int   Timeout for each request in seconds. (default 10)
```

#### Examples
Bruteforce all serverless resources from an AWS profile
```
❯❯ cliam aws serverless --profile=my-profile
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
Currently, permissions are enumerate using the `cloudresourcemanager` API. This will fail if this service is not enabled, but there are future plans to extend using rest alls to confirm permissions.

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

## Debug
cliam supports two environment variables to show debug output
- DEBUG=true (shows status codes of requests)
- VERBOSE=true (shows body of requests)
