param(
    [Parameter(Mandatory=$false)]
    [string]$OutputPath = "keystone_compilation_fix.patch"
)

$SCRIPT_DIR = Split-Path -Parent $MyInvocation.MyCommand.Path
$ErrorActionPreference = "Stop"

Write-Host "=== HyperDbg Patch Generator ===" -ForegroundColor Cyan

$keystoneDir = Join-Path $SCRIPT_DIR "clone/keystone"
if (-not (Test-Path $keystoneDir)) { Write-Error "keystone dir not found"; exit 1 }

Push-Location $keystoneDir
try {
    $patchPath = Join-Path $SCRIPT_DIR $OutputPath
    Remove-Item $patchPath -Force -ErrorAction SilentlyContinue

    git diff --output="$patchPath" 2>$null
    if (Test-Path $patchPath) {
        $size = (Get-Item $patchPath).Length
        if ($size -gt 0) { Write-Host "Patch: $patchPath ($size bytes)" -ForegroundColor Green }
        else { Remove-Item $patchPath; Write-Host "No changes" -ForegroundColor Yellow }
    } else { Write-Host "No changes" -ForegroundColor Yellow }
} finally { Pop-Location }
