# script\build.ps1: Script to build the binary

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

# Change working directory to script's grandparents directory
Set-Location -Path (Get-Item (Split-Path -Parent $MyInvocation.MyCommand.Definition)).Parent.FullName

.\script\clean.ps1

# Get Git SHA and Version
$GIT_SHA = git rev-parse --short HEAD
$VERSION = git rev-parse --abbrev-ref HEAD
$LDFLAGS = "-X github.com/dremio/statcapn/pkg/versions.gitSha=$GIT_SHA -X github.com/dremio/statcapn/pkg/versions.version=$VERSION"

# Build again and copy default-ddc.yaml
go build -ldflags "$LDFLAGS" -o .\bin\statcapn.exe