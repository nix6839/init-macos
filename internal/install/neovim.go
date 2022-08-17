package install

import "github.com/nix6839/macos-init/pkg/command"

func Neovim() {
	command.New("nvim").
		Arg("+PlugInstall").
		Arg("+qall").
		Run()
}
