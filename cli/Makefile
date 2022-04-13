.PHONY: dev
.EXPORT_ALL_VARIABLES:

VERSION = dev

dev:
	go build -ldflags="-s -w -X github.com/securisec/iam-enumerate/version.BuildDate=2021 -X github.com/securisec/iam-enumerate/version.GitCommit=${COMMIT_HASH} -X github.com/securisec/iam-enumerate/version.GitBranch=${COMMIT_BRANCH} -X github.com/securisec/iam-enumerate/version.Version=${VERSION}" -o cliam