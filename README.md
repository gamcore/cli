<p align="center" style="align-content: center; text-align: center;">
<img src="manifest/goo.png" width="256" alt="logo" /><br />
</p>

# Goo

New Cross-Platform command line installer

## Installing

### Windows

#### Using `powershell`

```ps1
iex (New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/goo-app/cli/main/install.ps1')
```

#### Using `cmd`

```cmd
powershell -Command "iex (New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/goo-app/cli/main/install.ps1')"
```

### Linux (using `bash`) / MacOS (10.14+) (using `zsh`)

```sh
curl -s https://raw.githubusercontent.com/goo-app/cli/main/install.sh | bash
```

## Usage

### manage app installation

```shell
goo install <app...>
goo update [-c|--cleanup] <app...>
goo uninstall <app...>
```

### listing of installed apps

```shell
goo list
```

### searching available app

```shell
goo search [--regex] <app>
```

### manage repository

```shell
goo repo add [--name <name>] <git-url>
goo repo remove <name>
```
