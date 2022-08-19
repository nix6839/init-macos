package install

import "github.com/nix6839/macos-init/pkg/command"

var masPackages = []string{
	// KakaoTalk
	"869223134",
}

func Mas() {
	command.New("mas").Arg("install").Args(masPackages).Run()
}
