#!/bin/bash

GOO_PATH="${HOME}/.goo"

abort() {
  printf "%s \n" "$@"
  exit 1
}

MkDirIfNotExist() {
  for d in "$@"; do
    if [ ! -d "$d" ]; then
      mkdir -p "$d"
    fi
  done
}

Download() {
  GOO_DOWNLOAD_URL="$(curl -H "Accept: application/vnd.github.v3+json" -s https://api.github.com/repos/stachu540/goo/releases/latest | grep "browser_download_url.*${GOO_NAME}" | cut -d : -f 2,3 | tr -d \")"
  curl -o "${GOO_PATH}/apps/goo/current/goo" "$GOO_DOWNLOAD_URL"
}

Prepare() {
  if [ -d "$GOO_PATH" ]; then
    abort "Goo is already installed"
  fi
  MkDirIfNotExist "$GOO_PATH" "$GOO_PATH/bin" "$GOO_PATH/apps/goo/current" "$GOO_PATH/tmp" "$GOO_PATH/repos"
}

Install() {
  Download
  cd "${HOME}/.goo/bin" || ln -s "../apps/goo/current/goo" goo
  PATH="${HOME}/.goo/bin:$PATH"
  goo init
}

InstallLinux() {
  Prepare
  UNAME_MACHINE="$(uname -m)"

  if [[ "${UNAME_MACHINE}" == "arm64" ]]; then
    GOO_NAME="goo_linux_arm64"
  elif [[ "${UNAME_MACHINE}" == "x86_64" ]]; then
    GOO_NAME="goo_linux_amd64"
  elif [[ "${UNAME_MACHINE}" == "i386" ]]; then
    GOO_NAME="goo_linux_386"
  else
    abort "Unsupported architecture: ${UNAME_MACHINE}"
  fi
  Install
  echo "export GOO_PATH=\${HOME}/.goo" >> ~/.bashrc
  echo "export PATH=\${GOO_PATH}/bin:\$PATH" >> ~/.bashrc
}

InstallDarwin() {
  Prepare
  UNAME_MACHINE="$(uname -m)"

  if [[ "${UNAME_MACHINE}" == "arm64" ]]; then
    GOO_NAME="goo_darwin_arm64"
  else
    GOO_NAME="goo_darwin_amd64"
  fi
  Install
  echo "export GOO_PATH=\${HOME}/.goo" >> ~/.zshrc
  echo "export PATH=\${GOO_PATH}/bin:\$PATH" >> ~/.zshrc
}

if [ -z "${BASH_VERSION:-}" ]; then
  abort "Bash is required to interpret this script"
fi

if [[ ! -t 0 || -n "${CI-}" ]]
then
  NONINTERACTIVE=1
fi

OS="$(uname)"

if [[ "${OS}" == "Linux" ]]; then
  InstallLinux
elif [[ "${OS}" == "Darwin" ]]; then
  InstallDarwin
else
  abort "Goo supports only macOS, Linux and Windows!"
fi
