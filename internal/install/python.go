package install

import (
	"github.com/nix6839/macos-init/pkg/command"
	"github.com/nix6839/macos-init/pkg/utils"
)

func fetchPoetryScript() string {
	return utils.FetchBody("https://install.python-poetry.org")
}

func Python() {
	command.New("python3").
		Arg("-c").
		Arg(fetchPoetryScript()).
		Run()
}
