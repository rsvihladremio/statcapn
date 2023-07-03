# Script to build binaries in all supported platforms

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

# Change working directory to script's grandparents directory
Set-Location -Path (Get-Item (Split-Path -Parent $MyInvocation.MyCommand.Definition)).Parent.FullName

# Get Git SHA and Version
$GIT_SHA = git rev-parse --short HEAD
$VERSION = $args[0]
$LDFLAGS = "-X github.com/rsvihladremio/statcapn/pkg/versions.gitSha=$GIT_SHA -X github.com/rsvihladremio/statcapn/pkg/versions.version=$VERSION"

Write-Output "Cleaning bin folder…"
Get-Date -Format "HH:mm:ss"
.\script\clean

Write-Output "Building linux-amd64…"
Get-Date -Format "HH:mm:ss"
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -ldflags "$LDFLAGS" -o ./bin/statcapn
Compress-Archive -Path ./bin/statcapn -DestinationPath ./bin/statcapn-linux-amd64.zip

Write-Output "Building linux-arm64…"
Get-Date -Format "HH:mm:ss"
$env:GOARCH="arm64"
go build -ldflags "$LDFLAGS" -o ./bin/statcapn
Compress-Archive -Path ./bin/statcapn -DestinationPath ./bin/statcapn-linux-arm64.zip

Write-Output "Building darwin-os-x-amd64…"
Get-Date -Format "HH:mm:ss"
$env:GOOS="darwin"
$env:GOARCH="amd64"
go build -ldflags "$LDFLAGS" -o ./bin/statcapn
Compress-Archive -Path ./bin/statcapn  -DestinationPath ./bin/statcapn-darwin-amd64.zip

Write-Output "Building darwin-os-x-arm64…"
Get-Date -Format "HH:mm:ss"
$env:GOARCH="arm64"
go build -ldflags "$LDFLAGS" -o ./bin/statcapn
Compress-Archive -Path ./bin/statcapn -DestinationPath ./bin/statcapn-darwin-arm64.zip

Write-Output "Building windows-amd64…"
Get-Date -Format "HH:mm:ss"
$env:GOOS="windows"
$env:GOARCH="amd64"
go build -ldflags "$LDFLAGS" -o ./bin/statcapn.exe
Compress-Archive -Path ./bin/statcapn.exe -DestinationPath ./bin/statcapn-windows-amd64.zip

Write-Output "Building windows-arm64…"
Get-Date -Format "HH:mm:ss"
$env:GOOS="windows"
$env:GOARCH="arm64"
go build -ldflags "$LDFLAGS" -o ./bin/statcapn.exe
Compress-Archive -Path ./bin/statcapn.exe  -DestinationPath ./bin/statcapn-windows-arm64.zip
