$scriptPath = Split-Path -Parent $MyInvocation.MyCommand.Path
$patchPath = Join-Path $scriptPath "driver_compilation_fix.patch"
$hyperdbgPath = Join-Path $scriptPath "HyperDbg"

if (-not (Test-Path $patchPath)) {
    Write-Error "Patch file not found: $patchPath"
    exit 1
}

if (-not (Test-Path $hyperdbgPath)) {
    Write-Error "HyperDbg directory not found: $hyperdbgPath"
    exit 1
}

Push-Location $hyperdbgPath
try {
    git apply $patchPath
    Pop-Location
    git status --short
} catch {
    Pop-Location
    Write-Error "Failed to apply patch: $_"
    exit 1
}
