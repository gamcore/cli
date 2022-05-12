$GooPath = ""

### DO NOT EDIT BELOW THIS LINE

function abort {
  param($message)
  Write-Error -Message $message
  exit 1
}

function MkDirIfNotExist([string[]]$paths) {
  foreach ($path in $paths) {
    if (!(Test-Path $path)) {
      New-Item -ItemType Directory -Path $path | Out-Null
    }
  }
}

function Download {
  $latest = iwr -UseBasicParsing -Headers @{"Accept"="application/vnd.github.v3+json"} -Uri "https://api.github.com/repos/goo-app/cli/releases/latest"
  $file = $($($latest | ConvertFrom-Json).assets | where { $_.name -match "^$GOO_NAME$" }).browser_download_url
  Invoke-WebRequest -Uri $file -OutFile "$GOO_PATH/apps/goo/current/goo.exe"
}

function Prepare {
  if (Test-Path $GOO_PATH) {
    abort "Goo is already installed"
  }
  MkDirIfNotExist(@($GOO_PATH,"$GOO_PATH","$GOO_PATH\bin","$GOO_PATH\apps\goo\current","$GOO_PATH\tmp","$GOO_PATH\repos"))
}

function Install {
  Download
  Set-Location -Path "$GOO_PATH/bin"
  New-Item -ItemType SymbolicLink -Path "$GOO_PATH\bin\goo.exe" -Value "..\apps\goo\current\goo.exe"
}

function InstallWindows {
  Prepare
  $WindowsArchitecture = $(gwmi -Query "Select OSArchitecture from Win32_OperationgSystem").OSArchitecture

  $arch = if ($WindowsArchitecture.Contains("ARM")) {
    if ($WindowsArchitecture.StartsWith("64")) {
      "arm64"
    } else {
      "arm"
    }
  } else {
    if ($WindowsArchitecture.StartsWith("64")) {
      "amd64"
    } else {
      "386"
    }
  }

  $GOO_NAME="goo-windows-$arch.exe"
  Install
  [Environment]::SetEnvironmentVariable("GOO_PATH", $GOO_PATH, "User")
  [Environment]::SetEnvironmentVariable("PATH", "$env:PATH;%GOO_PATH%\bin", "User")
}

function InstallOtherOS {
  Invoke-Expression -Command "bash <(curl -s https://raw.githubusercontent.com/goo-app/cli/main/install.sh)"
}

$GOO_PATH = if (($GooPath -eq "") -or ($GooPath -eq $null)) {
  if (($env:GOO_PATH -eq $null) -or ($env:GOO_PATH -eq "")) {
    "$env:USERPROFILE/.goo"
  } else {
    $env:GOO_PATH
  }
} else {
  $GooPath
}

if ($IsWindows) {
  InstallWindows
} elseif ($IsLinux -or $IsMacOS) {
  InstallOtherOS
} else {
  Write-Error "Goo supports only macOS, Linux and Windows"
  exit 1
}
