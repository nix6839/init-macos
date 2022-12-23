package prepare

import (
	"github.com/nix6839/macos-init/pkg/command"
	"github.com/nix6839/macos-init/pkg/utils"
)

func fetchBrewScript() string {
	return utils.FetchBody("https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh")
}

var (
	fontPackages = []string{
		"font-d2coding", "font-jetbrains-mono", "font-jetbrains-mono-nerd-font",
		"font-pretendard",
	}

	cliPackages = []string{
		"git", "neovim", "chezmoi", "lsd", "bat", "tealdeer", "git-delta",
		"gnupg", "pinentry-mac", "shellcheck", "go", "golangci-lint", "mas",
		"python@3.10", "mkcert",
	}

	caskPackages = []string{
		"gureumkim", "microsoft-edge", "alacritty", "discord", "bitwarden",
		"visual-studio-code", "notion", "slack", "spotify", "telegram-desktop",
		"dropbox", "remix-ide", "firefox", "figma",
		// Does not support m1 native
		"tutanota",
	}

	caskNoQuarantinePackages = []string{
		"chromium",
	}
)

func Brew() {
	command.New("/bin/bash").
		Arg("-c").
		Arg(fetchBrewScript()).
		Run()

	command.New("brew").
		Arg("tap").
		Arg("homebrew/cask-fonts").
		Run()

	command.New("brew").
		Arg("install").
		Args(fontPackages).Args(cliPackages).Args(caskPackages).
		Run()

	command.New("brew").
		Arg("install").
		Arg("--no-quarantine").
		Args(caskNoQuarantinePackages).
		Run()
}
