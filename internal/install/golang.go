package install

import "github.com/nix6839/macos-init/pkg/command"

func Golang() {
	command.New("go").
		Arg("install").
		Arg("golang.org/x/tools/gopls@latest").
		Run()
}
