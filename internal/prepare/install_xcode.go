package prepare

import (
	"github.com/nix6839/macos-init/pkg/command"
)

func InstallXcode() {
	command.New("xcode-select").Arg("--install").Run()
}
