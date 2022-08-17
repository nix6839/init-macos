package main

import (
	"github.com/nix6839/macos-init/internal/install"
	"github.com/nix6839/macos-init/internal/prepare"
)

func main() {
	prepare.SetZshenv()
	prepare.InstallXcode()
	prepare.Brew()
	prepare.Chezmoi()

	install.Zsh()
	install.Neovim()
	install.Nvm()
	install.Rust()
	install.Golang()
	install.VSCode()
}
