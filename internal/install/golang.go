package install

import "github.com/nix6839/macos-init/pkg/command"

var golangPackages = []string{
	"golang.org/x/tools/gopls@latest",
}

func Golang() {
	for _, golangPackage := range golangPackages {
		command.New("go").
			Arg("install").
			Arg(golangPackage).
			Run()
	}
}
