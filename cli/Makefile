.PHONY: dev
.EXPORT_ALL_VARIABLES:

VERSION = dev

dev:
	go build -ldflags="-s -w -X github.com/securisec/cliam/cli/version.BuildDate=2022 -X github.com/securisec/cliam/cli/version.GitCommit=${COMMIT_HASH} -X github.com/securisec/cliam/cli/version.GitBranch=${COMMIT_BRANCH} -X github.com/securisec/cliam/cli/version.Version=${VERSION}" -o cliam