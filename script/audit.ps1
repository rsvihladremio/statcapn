# script\audit.ps1: runs gosec against the mod file to find security issues

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

# Change working directory to script's grandparents directory
Set-Location -Path (Get-Item (Split-Path -Parent $MyInvocation.MyCommand.Definition)).Parent.FullName

Write-Output "Running gosec..."
Invoke-Expression "gosec  ./..."
