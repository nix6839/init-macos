package prepare

import (
	"github.com/nix6839/macos-init/pkg/command"
	"github.com/nix6839/macos-init/pkg/utils"
)

func fetchChezmoiScript() string {
	return utils.FetchBody("https://www.chezmoi.io/get")
}

func Chezmoi() {
	command.New("sh").
		Arg("-c").
		Arg(fetchChezmoiScript()).
		Arg("--").
		Arg("init").
		Arg("--apply").
		Arg("https://github.com/nix6839/dotfiles-macos.git").
		Run()
}
